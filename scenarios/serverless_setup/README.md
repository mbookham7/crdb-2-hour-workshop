# Setting Up a Serverless Cluster

During this section of the lab you will learn how to use Cockroach Cloud to create a CockroachDB Serverless cluster. Once this is up and running, which only takes a few seconds, you will then learn how to connect to this with code and the CLI client.

To kick thinks off lets create our serverless cluster

## Setup Your Serverless Cluster

In your browser access the following URL
```bash
https://cockroachlabs.cloud/
```

Login to the console, in the centre of the screen click `Create Cluster`

![create-cluster](/images/serverless-setup/create-cluster.png)

Once you have clicked on `Create Cluster` you will be presented with the a wizard that walks you through the setup of your serverless cluster.

In this scenario we will be creating we will be creating a multi regional serverless cluster in AWS. Under the Cloud Provider pick AWS.

![select-your-cloud](/images/serverless-setup/select-your-cloud.png)

Next, we are going to select the regions. As we are going to deploy a multi-region serverless cluster we will need to add two additional regions. Please select the following regions for this scenario.

- N. Virginia (us-east-1)
- Oregon (us-west-2)
- Frankfurt (eu-central-1)

![select-your-regions](/images/serverless-setup/select-your-regions.png)

Your first cluster can be free if you stay with in the limits which we will within this lab. You must enter a unique name for your cluster now user you name to make it unique. 

`yourname-serverless-lab`

Enter in the box below.

![name-your-cluster](/images/serverless-setup/name-your-cluster.png)

Once you have made your selection review them and hit create!! WE ARE GOING GLOBAL!! Once your cluster is created then you will need to create an SQL User. The user name will be based on the account you logged in with but the password will be randomly generated, you need to copy this to you notepad and keep this safe for further steps in the lab.

![create-sql-user](/images/serverless-setup/create-sql-user.png)

## Connecting to your Cluster via Code

@Rob Reid

## Connecting to your Cluster via CLI

To connect to your cluster you will first need the `CockraochDB Client`. You can do this from the console by clicking connect in the top right hand corner.

Below are the commands for Mac and Windows. Copy the correct one for your Operating System and run.

MacOS
```shell
curl https://binaries.cockroachdb.com/cockroach-v23.1.10.darwin-10.9-amd64.tgz | tar -xz; sudo cp -i cockroach-v23.1.10.darwin-10.9-amd64/cockroach /usr/local/bin/
```

Windows Powershell
```powershell
$ErrorActionPreference = "Stop"; [Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12; $ProgressPreference = 'SilentlyContinue'; $null = New-Item -Type Directory -Force $env:appdata/cockroach; Invoke-WebRequest -Uri https://binaries.cockroachdb.com/cockroach-v23.1.10.windows-6.2-amd64.zip -OutFile cockroach.zip; Expand-Archive -Force -Path cockroach.zip; Copy-Item -Force cockroach/cockroach-v23.1.10.windows-6.2-amd64/cockroach.exe -Destination $env:appdata/cockroach; $Env:PATH += ";$env:appdata/cockroach"; # We recommend adding ";$env:appdata/cockroach" to the Path variable for your system environment. See https://docs.microsoft.com/en-us/powershell/module/microsoft.powershell.core/about/about_environment_variables#saving-changes-to-environment-variables for more information.
```

Once you have downloaded the CockroachDB Client you are now ready to connect to your cluster. From the same window, from the connect button in the top right hand corner you can copy your connection command.

Example command below
```shell
cockroach sql --url "postgresql://<username>@<cluster-name>.cockroachlabs.cloud:26257/defaultdb?sslmode=verify-full"
```

Now we are connected to the cluster we are going to run some basic SQL commands to get familiar with COckroachDB.

``

@Rob Reid






