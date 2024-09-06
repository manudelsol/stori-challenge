# CSV Processing Project for Stori Challenge

## Overview & Features

This project implements a Lambda function (`storiLambda`) that:

1. Takes a CSV file from an S3 bucket containing account transactions.
2. Processes the transactions by separating them by month and type (credit or debit).
3. Stores account data and each transaction into a PostgreSQL database.
4. Sends a summary email to the specified recipient using Mailjet.

### CSV Format

The CSV file should follow this format:

```
Id,Date,Transaction
0,3/6,+120
1,3/11,-104.5
2,3/20,+54.5
3,4/15,-30
4,5/1,+28
5,5/25,-54.3
6,6/20,+40
```

### Sample JSON Input

The Lambda function accepts the following JSON structure as input:

```json
{
  "bucket": "your-bucket-name",
  "key": "csv-file-key",
  "email": "recipient-email"
}
```

### Database Schema

The project uses PostgreSQL to store account and transaction information. The tables are structured as follows:

#### `accounts` Table

| Column        | Data Type   | Description                         |
| ------------- | ----------- | ----------------------------------- |
| id            | bigserial   | Primary key (auto-incremented).     |
| created_at    | timestamptz | Timestamp when the account is added.|
| user_email    | varchar     | Email associated with the account.  |
| balance       | float       | Current balance of the account.     |
| number_of_txs | int         | Number of transactions processed.   |

#### `txns` Table

| Column      | Data Type | Description                           |
| ----------- | --------- | ------------------------------------- |
| id          | bigserial | Primary key (auto-incremented).       |
| created_at  | date      | Date of the transaction.              |
| tx          | varchar   | Transaction details (credit or debit).|
| account_id  | int       | Foreign key referencing the `accounts` table. |

---

## Setup Instructions

### Prerequisites

- **Go 1.18+**
- **PostgreSQL 12+**
- **AWS S3** for CSV storage
- **AWS Lambda** for function deployment
- **Mailjet API** for sending emails

### Environment Variables

Set the following environment variables for local development and deployment:

- `DB_HOST`: PostgreSQL host address.
- `DB_PORT`: PostgreSQL port.
- `DB_USER`: PostgreSQL username.
- `DB_PASSWORD`: PostgreSQL password.
- `DB_NAME`: PostgreSQL database name.
- `MAILJET_API_KEY`: Mailjet API key.
- `MAILJET_SECRET_KEY`: Mailjet secret key.
- `MAILJET_TEMPLATE_ID`: The ID of the email template.

### Installing Dependencies

Run the following command to install project dependencies:

```bash
go mod tidy
```

---

## AWS Lambda Deployment

1. **Build the Lambda Binary**: Use the following command to build your Go binary for AWS Lambda:

```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go
```

2. **Zip the Binary**:

```bash
zip function.zip main
```

3. **Deploy to Lambda**:
   - Upload the `function.zip` file to your Lambda function.
   - Set the appropriate environment variables in the Lambda console.

### IAM Permissions

Ensure your Lambda function has the following permissions:

- **S3 Read Access**: To fetch the CSV file.
- **PostgreSQL Access**: For database connectivity (if using AWS RDS, ensure the Lambda has access to the RDS instance).
- **SES or Mailjet API Access**: To send email notifications.

---

## AWS Lambda Invocation from CLI

To invoke the Lambda function from the AWS CLI, use the following command:

```bash
aws lambda invoke \
    --function-name functionName \
    --cli-binary-format raw-in-base64-out \
    --payload '{"bucket": "bucketName", "key": "keyOfCsvFile", "email": "recipientEmail"}' \
    response.json
```

- Replace `functionName` with the name of your Lambda function (e.g., `storiLambda`).
- Replace `bucketName` with the name of the S3 bucket where your CSV file is stored.
- Replace `keyOfCsvFile` with the key or path of the CSV file inside the bucket.
- Replace `recipientEmail` with the email address where you want to send the summary.

The response of the Lambda invocation will be written to the `response.json` file.

---
