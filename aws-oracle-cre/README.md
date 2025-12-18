# AWS API Oracle Workflow

**Part of:** [Chainlink CRE + AWS Oracle Demo](../docs/README.md)

## Overview

CRE workflow that fetches asset price and proof of reserves data from AWS API, calculates collateralization ratios, validates thresholds, and writes results to Ethereum smart contracts.

### Key Features

- Parallel data fetching with DON consensus
- Collateralization ratio calculation and threshold validation
- Dual contract updates (PriceFeed + CollateralizationMonitor)
- Configurable cron schedule (default: every 10 seconds)

## Architecture

### CRE Workflow

```
Cron Trigger (*/10 * * * * *)
         ↓
┌────────────────────────────┐
│  Parallel HTTP Requests    │
│  • GET /asset-price        │
│  • GET /proof-of-reserves  │
└────────────────────────────┘
         ↓
Calculate Ratio & Validate
         ↓
┌────────────────────────────────┐
│  Sequential Contract Writes    │
│  1. PriceFeed                  │
│  2. CollateralizationMonitor   │
└────────────────────────────────┘
```

### Workflow Execution

1. Cron trigger executes workflow (configurable schedule)
2. Fetch price and reserves from AWS API with DON consensus
3. Calculate collateralization ratio (reserves / price)
4. Validate ratio against threshold (default: 1.2)
5. Write price to PriceFeed contract
6. Write collateral data to CollateralizationMonitor contract
7. Return result with transaction hashes

## Prerequisites

- **AWS Backend**: Deploy the AWS serverless backend first (see [../price-feed-por-dynamodb-crud/README.md](../price-feed-por-dynamodb-crud/README.md))
- **CRE CLI**: v1.0.3+ ([Installation Guide](https://docs.chain.link/cre/getting-started/cli-installation))
  - After installation, run `cre login` to authenticate
  - Verify with `cre whoami`
- **Go**: 1.25.3+ ([Install Go](https://go.dev/doc/install))
- **Sepolia ETH**: Get testnet ETH from [faucets.chain.link](https://faucets.chain.link)

## Quick Start

### 1. Setup Environment

```bash
cd ~/sample-cre-pricefeeds-por/aws-oracle-cre

# Run script to create .env with API key and update config with API URL
./set-api-key.sh

# Edit .env and add your Ethereum private key
# Replace: CRE_ETH_PRIVATE_KEY=your_64_char_private_key_without_0x_prefix
```

**Required environment variables in `.env`:**
- `API_KEY_VALUE` - AWS API Gateway key (auto-populated by script)
- `CRE_ETH_PRIVATE_KEY` - Your Ethereum private key (you must add this manually)
  - 64 hexadecimal characters
  - No `0x` prefix
  - Get from MetaMask: Account Details → Show Private Key

**Get testnet ETH:** [faucets.chain.link](https://faucets.chain.link)

### 2. Deploy Smart Contracts

Deploy the PriceFeed and CollateralizationMonitor contracts to Sepolia. See [contracts/evm/README.md](contracts/evm/README.md) for detailed instructions.

The deploy script automatically updates `config.staging.json` with the contract addresses.

### 3. Generate Contract Bindings

```bash
cd ~/sample-cre-pricefeeds-por/aws-oracle-cre
cre generate-bindings evm
go mod tidy
```

### 4. Run Simulation

**Option A: Using the automated script (recommended)**

> ⚠️ **Note:** This script broadcasts transactions to Sepolia testnet and will consume gas.

```bash
cd ~/sample-cre-pricefeeds-por/aws-oracle-cre
./run-simulation.sh
```

**Option B: Manual CLI commands**

Dry run (no transactions broadcast):
```bash
cre workflow simulate api-oracle --target staging-settings
```

With broadcast (sends real transactions to Sepolia):
```bash
cre workflow simulate api-oracle --target staging-settings --broadcast
```

> 💡 **Tip:** Use dry run first to verify your workflow compiles and runs correctly before broadcasting transactions.

## Configuration

### Workflow Settings (`api-oracle/workflow.yaml`)

Defines workflow name and paths for each target environment:

```yaml
staging-settings:
  user-workflow:
    workflow-name: "api-oracle-staging"
  workflow-artifacts:
    workflow-path: "."
    config-path: "./config.staging.json"
    secrets-path: "../secrets.yaml"
```

### Config Files

**`api-oracle/config.staging.json`** / **`config.production.json`**

| Parameter | Description | Example |
|-----------|-------------|---------|
| `schedule` | Cron expression | `*/10 * * * * *` (every 10 seconds) |
| `apiUrl` | AWS API Gateway URL | `https://{api-id}.execute-api.{region}.amazonaws.com/prod/` |
| `minCollateralizationRatio` | Minimum threshold | `1.2` (120%) |
| `evms[].chainSelector` | Chain ID | `16015286601757825753` (Sepolia) |
| `evms[].priceFeedAddress` | PriceFeed contract address | From deployment |
| `evms[].collateralizationMonitorAddress` | Monitor contract address | From deployment |
| `evms[].gasLimit` | Max gas per transaction | `500000` |

### Secrets Configuration

**`secrets.yaml`** - Maps logical secret IDs to environment variables:

```yaml
secretsNames:
    API_KEY:
        - API_KEY_VALUE
```

**`.env`** - Provides actual secret values:

```bash
API_KEY_VALUE=your_api_gateway_key
CRE_ETH_PRIVATE_KEY=your_64_char_private_key_without_0x_prefix
```

See [CRE Secrets Documentation](https://docs.chain.link/cre/guides/workflow/secrets/using-secrets-simulation) for details.

## Project Structure

```
aws-oracle-cre/
├── api-oracle/              # Workflow directory
│   ├── main.go             # Workflow implementation
│   ├── workflow.yaml       # Workflow settings
│   ├── config.staging.json # Staging configuration
│   └── config.production.json # Production configuration
├── contracts/evm/          # Smart contracts
├── project.yaml            # Project-level settings (RPC URLs)
├── secrets.yaml            # Secret mappings
├── .env                    # Secret values (not committed)
├── run-simulation.sh       # Quick simulation script
└── set-secret.sh          # AWS API key retrieval script
```

## Deployment

See [contracts/evm/README.md](contracts/evm/README.md) for smart contract deployment instructions.

For production deployment of the workflow to CRE DON, see [CRE Deployment Guide](https://docs.chain.link/cre/guides/operations/deploying-workflows).

## Monitoring

View transaction results on [Sepolia Etherscan](https://sepolia.etherscan.io/).

## Troubleshooting

### "Secret not found"
- Ensure both `API_KEY_VALUE` and `CRE_ETH_PRIVATE_KEY` are set in `.env`
- Check `secrets.yaml` maps `API_KEY` to `API_KEY_VALUE`
- Verify `.env` file exists in project root

### "403 Forbidden" from API
- Run `./set-api-key.sh` to get current API key from AWS
- Verify API Gateway is deployed and accessible
- Check API key is valid in AWS Console

### "Invalid private key" or "invalid hex character"
- Ensure `CRE_ETH_PRIVATE_KEY` is exactly 64 hexadecimal characters (0-9, a-f)
- Remove `0x` prefix if present
- No spaces, quotes, or special characters
- Example format: `abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890`

### "Failed to parse private key"
- Your private key must be the raw 64-character hex string
- Export from MetaMask: Account Details → Show Private Key
- Remove the `0x` prefix before adding to `.env`

## Documentation

- [CRE Documentation](https://docs.chain.link/cre)
- [Smart Contract Guide](./contracts/evm/README.md)

## License

ISC
