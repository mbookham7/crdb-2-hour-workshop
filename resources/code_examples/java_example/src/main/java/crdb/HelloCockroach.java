package crdb;

import java.sql.*;

public class HelloCockroach {
  public static void main(String[] args) throws SQLException {
    String connectionString = System.getenv("CONNECTION_STRING");

    try (Connection conn = DriverManager.getConnection(connectionString)) {
      try (Statement st = conn.createStatement()) {
        try (ResultSet rs = st.executeQuery("SELECT id, full_name FROM member")) {
          while (rs.next()) {
            System.out.printf(
                "%s: %s\n",
                rs.getString(1),
                rs.getString(2));
          }
          rs.close();
          st.close();
        }
      }
    }
  }
}