using System;
using System.Data;
using System.Net.Security;
using Npgsql;

var connectionString = Environment.GetEnvironmentVariable("CONNECTION_STRING");

using (var conn = new NpgsqlConnection(connectionString))
{
  conn.Open();

  using (var cmd = new NpgsqlCommand("SELECT id, full_name FROM member", conn))
    using (var reader = cmd.ExecuteReader())
      while (reader.Read())
        Console.WriteLine("{0}: {1}", reader.GetValue(0), reader.GetValue(1));
}