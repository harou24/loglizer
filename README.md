# loglizer
Loglizer is a web-based interface for summarizing log files
capable of processing large amounts of data by utilizing a multi-stage pipeline consisting of 
reader, worker, and combiner stages to parallelize and streamline the analysis.

# Running the Server
To start the server, open your terminal and run the following command from the root directory of the project:

```azure
make run
```
This command will start the server, making it listen for incoming requests. Ensure you have make installed and that you're in the correct directory where the Makefile is located.

# Testing the Server
Once the server is up and running, you can test its functionality by sending a POST request with a CSV file. Replace /path/to/journaux.csv with the actual path to your CSV file that you want to analyze.

Use the following curl command to send your request:

```azure
curl -X POST -H "Accept: text/csv" -F "file=@/path/to/journaux.csv" http://localhost:15442/analysis
```
