# Asset Price and Reserves Service

Serverless REST API for tracking asset prices and proof of reserves data using AWS Lambda, API Gateway, and DynamoDB.

**Part of:** [Chainlink CRE + AWS Oracle Demo](../docs/README.md)

## Quick Start

```bash
cd ~/sample-cre-pricefeeds-por/price-feed-por-dynamodb-crud
./deploy.sh
source ./setup-env.sh
./test-api.sh
```

**Next Step:** [Deploy CRE Workflow](../aws-oracle-cre/README.md) to consume this API.

## Architecture

```
API Gateway → Lambda Functions → DynamoDB
```

- **API Gateway**: REST API with API key authentication
- **Lambda Functions**: 5 serverless handlers (TypeScript/Node.js 20.x)
- **DynamoDB**: Single table with composite key (`recordType`, `timestamp`)

## Prerequisites

- Node.js 20.x or later
- AWS CLI configured with appropriate credentials
- AWS SAM CLI installed
- esbuild installed globally: `npm install -g esbuild`

## Deployment

Deploy the stack:
```bash
cd ~/sample-cre-pricefeeds-por/price-feed-por-dynamodb-crud
./deploy.sh
```

Verify deployment and set environment variables:
```bash
./verify-deployment.sh
source ./setup-env.sh
```

Test the API:
```bash
./test-api.sh
```

## API Endpoints

All endpoints require `x-api-key` header.

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/asset-price` | Store asset price |
| GET | `/asset-price/latest` | Get latest price |
| POST | `/proof-of-reserves` | Store proof of reserves |
| GET | `/proof-of-reserves/latest` | Get latest reserves |
| POST | `/simulate` | Generate test data |

## Example Usage

Generate test data:
```bash
curl -X POST "${API_URL}simulate" -H "x-api-key: ${API_KEY}"
```

Store asset price:
```bash
curl -X POST "${API_URL}asset-price" \
  -H "x-api-key: ${API_KEY}" \
  -H "Content-Type: application/json" \
  -d '{"price": 1234.56}'
```

Get latest price:
```bash
curl -X GET "${API_URL}asset-price/latest" -H "x-api-key: ${API_KEY}"
```

Response format:
```json
{"price": 1234.56, "timestamp": "2025-10-15T10:30:00.000Z"}
```

## Configuration

### Environment Variables

Set after deployment using `source ./setup-env.sh`:
- `API_URL`: API Gateway endpoint
- `API_KEY`: API key for authentication

### CloudFormation Parameters

Defined in `template.yaml`:
- Stack name: `asset-price-service`
- DynamoDB table: `AssetData`
- API Gateway stage: `Prod`

## Monitoring

View logs for all functions:
```bash
npm run logs
```

View logs for specific function:
```bash
sam logs -n StoreAssetPriceFunction --stack-name asset-price-service --tail
```

View stack resources:
```bash
aws cloudformation describe-stack-resources --stack-name asset-price-service
```

Key metrics to monitor:
- Lambda invocation count and errors
- DynamoDB read/write capacity usage
- API Gateway 4xx/5xx errors

## Troubleshooting

**Deployment fails with "Unable to upload artifact"**
- Ensure AWS CLI is configured with correct credentials
- Check IAM permissions for CloudFormation, Lambda, API Gateway, DynamoDB

**API returns 403 Forbidden**
- Verify API key is set: `echo $API_KEY`
- Check `x-api-key` header is included in request

**Lambda function timeout**
- Check CloudWatch logs: `npm run logs`
- Verify DynamoDB table exists and is accessible
- Check Lambda execution role has DynamoDB permissions

**No data returned from GET endpoints**
- Ensure data has been stored using POST endpoints
- Use `/simulate` endpoint to generate test data
- Check DynamoDB table for records in AWS Console

**Environment variables not set**
- Run `source ./setup-env.sh` (not `./setup-env.sh`)
- Verify deployment completed successfully
- Check CloudFormation outputs: `aws cloudformation describe-stacks --stack-name asset-price-service`

## Cost

AWS Free Tier (typical development usage): **$0/month**
- DynamoDB: 25 GB storage, 25 WCU, 25 RCU
- Lambda: 1M requests/month, 400,000 GB-seconds
- API Gateway: 1M requests/month (free for 12 months)

Beyond free tier: **~$1-5/month** for low-traffic applications

Cost breakdown:
- DynamoDB: $0.25/GB/month + $0.00065 per WCU + $0.00013 per RCU
- Lambda: $0.20 per 1M requests + $0.0000166667 per GB-second
- API Gateway: $3.50 per million requests

## Cleanup

Delete all resources:
```bash
npm run delete
```

Or manually:
```bash
sam delete --stack-name asset-price-service
```

## Related Documentation

- [Main Project Overview](../docs/README.md)
- [CRE Workflow Setup](../aws-oracle-cre/README.md) - Consumes this API
  
## License

MIT
