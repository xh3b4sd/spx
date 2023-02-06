# spx

Public data collector for the S&P500 based on the https://finance.yahoo.com API.
A Github Action is scheduled to regularly update a CSV file hosted on S3. That
CSV file can be integrated via a Grafana CSV Data Source using the plugin
https://grafana.com/grafana/plugins/marcusolsson-csv-datasource.

![Grafana](/asset/grafana.png)
