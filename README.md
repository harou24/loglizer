# loglizer
Log Frequency Tracker: Summarizing hourly top logs from CSV files

# How to run
```azure
Make
```

# How to test once server is running
```azure
curl -X POST -H "Accept: text/csv" -F "file=@/path/to/journaux.csv" http://localhost:15442/analysis
```