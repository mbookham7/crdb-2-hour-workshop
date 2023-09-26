Open a CockroachDB shell to your Serverless cluster:

``` sh
cockroach sql --url "postgresql://USER@CLUSTER_NAME.aws-eu-central-1.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
```

Create a regional-aware database:

``` sql
CREATE DATABASE cap_workshop
  PRIMARY REGION "aws-eu-central-1"
  REGIONS "aws-us-east-1", "aws-us-west-2";

USE cap_workshop;
```

Create some tables:

``` sql
-- A regional table to store products. Products for markets near our Frankfurt region will be stored within that region, while products near our US regions will be stored there.
CREATE TABLE "product" (
  "name" STRING NOT NULL,
  "market" STRING NOT NULL,
  "crdb_region" CRDB_INTERNAL_REGION AS (
    CASE
      WHEN "market" IN ('uk', 'ie', 'de', 'fr') THEN 'aws-eu-central-1'
      WHEN "market" IN ('us', 'mx') THEN 'aws-us-east-1'
      ELSE 'aws-eu-central-1'
    END
  ) STORED,
  "amount" DECIMAL NOT NULL,
  "currency" STRING NOT NULL,

  PRIMARY KEY ("crdb_region", "name", "market"),
  INDEX ("market")
) LOCALITY REGIONAL BY ROW;

-- A global table to store translations for products. This data will be stored across all regions, resulting in quick reads for everyone, everywhere.
CREATE TABLE "i18n" (
  "word" STRING NOT NULL,
  "lang" STRING NOT NULL,
  "translation" STRING NOT NULL,

  PRIMARY KEY("word", "lang"),
  INDEX i18n_lang_idx (lang) STORING (translation)
) LOCALITY GLOBAL;
```

Insert some test data:

``` sql
INSERT INTO product ("name", "market", "amount", "currency") VALUES
  ('Americano', 'uk', '3.80', 'gbp'),
  ('Latte', 'uk', '4.15', 'gbp'),
  ('Cappuccino', 'uk', '4.15', 'gbp'),

  ('Americano', 'de', '4.40', 'eur'),
  ('Latte', 'de', '4.80', 'eur'),
  ('Cappuccino', 'de', '4.80', 'eur'),

  ('Americano', 'us', '4.80', 'usd'),
  ('Latte', 'us', '5.24', 'usd'),
  ('Cappuccino', 'us', '5.24', 'usd'),
  
  ('Americano', 'mx', '82.03', 'mxn'),
  ('Latte', 'mx', '89.59', 'mxn'),
  ('Cappuccino', 'mx', '89.59', 'mxn');

INSERT INTO i18n ("word", "lang", "translation") VALUES
  ('Americano', 'en', 'Americano'),
  ('Latte', 'en', 'Latte'),
  ('Cappuccino', 'en', 'Cappuccino'),

  ('Americano', 'de', 'Americano'),
  ('Latte', 'de', 'Latté'),
  ('Cappuccino', 'de', 'Cappuccino'),
  
  ('Americano', 'zh', '美式咖啡'),
  ('Latte', 'zh', '拿铁'),
  ('Cappuccino', 'zh', '卡布奇诺'),
  
  ('Americano', 'ja', 'アメリカーノ'),
  ('Latte', 'ja', 'カプチーノ'),
  ('Cappuccino', 'ja', 'ラテ');
```

Query for products to show cross-region latencies:

``` sql
SELECT i.translation, p.market, p.amount, p.currency
FROM product p
JOIN i18n i ON p.name = i.word
WHERE p.market = 'uk'
AND i.lang = 'en';

SELECT i.translation, p.market, p.amount, p.currency
FROM product p
JOIN i18n i ON p.name = i.word
WHERE p.market = 'us'
AND i.lang = 'ja';
```

Add super regions:

``` sql
SET enable_super_regions = 'on';
ALTER DATABASE "cap_workshop" ADD SUPER REGION us VALUES "aws-us-east-1", "aws-us-west-2";
ALTER DATABASE "cap_workshop" ADD SUPER REGION eu VALUES "aws-eu-central-1";

SET override_multi_region_zone_config = true;
ALTER DATABASE "cap_workshop" CONFIGURE ZONE USING
  num_replicas = 3,
  constraints = '[]';
```

Debug statements:

``` sql
-- Get the replica ids for each of the region/azs.
SELECT DISTINCT
  split_part(split_part(unnest(replica_localities), ',', 1), '=', 2) region,
  split_part(split_part(unnest(replica_localities), ',', 2), '=', 2) az,
  unnest(replicas) replica
FROM [SHOW RANGES FROM TABLE product];

-- Show ranges by region.
WITH
  replicas AS (
    SELECT DISTINCT
      split_part(unnest(replica_localities), ',', 1) region,
      replicas
    FROM [SHOW RANGES FROM TABLE product]
  ),
  ranges AS (
    SELECT
      replicas,
      range_id
    FROM [SHOW RANGES FROM TABLE product]
  )
SELECT
  split_part(re.region, '=', 2) region,
  re.replicas,
  array_agg(ra.range_id) range_ids
FROM replicas re
JOIN ranges ra ON re.replicas = ra.replicas
GROUP BY re.region, re.replicas
ORDER BY region, replicas;
```