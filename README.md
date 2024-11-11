# Fetch Assessment
Backend assessment for fetch

# Set up and run details
In order to run the service (with Go installed) locally simply,
- use the command **go build** followed by **fetch-assessment.exe**
- navigate to **localhost:8080** with your browser or via Postman

# Routes
- **/receipts/process**
    - Method: POST 
    - Payload: Receipt JSON
    - Response: JSON containing the id for the receipt
- **/receipts/{id}/points**
    - Method: GET
    - Response: JSON object containing the number of points belonging to the receipt with the given id provided

# Considerations
- Future iterations could have unit testing added
- Storing data in memory in production would not be ideal if data is required to be non-volatile, switching to a DB woudl be a logical follow-up in that case
- Validation is handled at a type level, so ensuring no malicious actors / payloads was considered here. 