#!/bin/bash

# Asset Price Service Deployment Script
# This script builds, packages, and deploys the Asset Price Service to AWS

set -e  # Exit on any error

echo "=========================================="
echo "Asset Price Service Deployment"
echo "=========================================="
echo ""

# Check if AWS CLI is installed
if ! command -v aws &> /dev/null; then
    echo "❌ Error: AWS CLI is not installed"
    echo "Please install AWS CLI: https://aws.amazon.com/cli/"
    exit 1
fi

# Check if SAM CLI is installed
if ! command -v sam &> /dev/null; then
    echo "❌ Error: AWS SAM CLI is not installed"
    echo "Please install SAM CLI: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/install-sam-cli.html"
    exit 1
fi

# Check AWS credentials
echo "🔍 Checking AWS credentials..."
if ! aws sts get-caller-identity &> /dev/null; then
    echo "❌ Error: AWS credentials not configured"
    echo "Please configure AWS credentials using 'aws configure'"
    exit 1
fi

AWS_ACCOUNT=$(aws sts get-caller-identity --query Account --output text)
AWS_REGION=$(aws configure get region || echo "us-east-1")
echo "✅ AWS Account: $AWS_ACCOUNT"
echo "✅ AWS Region: $AWS_REGION"
echo ""

# Step 1: Install dependencies
echo "📦 Step 1: Installing dependencies..."
npm install
echo "✅ Dependencies installed"
echo ""

# Step 2: Build SAM application (esbuild handles TypeScript compilation)
echo "📦 Step 2: Building SAM application with esbuild..."
sam build
echo "✅ SAM application built"
echo ""

# Step 3: Deploy to AWS
echo "🚀 Step 3: Deploying to AWS..."
echo "This may take a few minutes..."
sam deploy --no-confirm-changeset --no-fail-on-empty-changeset

if [ $? -eq 0 ]; then
    echo ""
    echo "=========================================="
    echo "✅ Deployment Successful!"
    echo "=========================================="

    echo ""
    echo "🔑 To get your API Key value, run:"
    API_KEY_ID=$(aws cloudformation describe-stacks \
        --stack-name asset-price-service \
        --query 'Stacks[0].Outputs[?OutputKey==`ApiKeyId`].OutputValue' \
        --output text)
    echo "aws apigateway get-api-key --api-key $API_KEY_ID --include-value --query 'value' --output text"
    echo ""
    
    # Get API endpoint
    API_URL=$(aws cloudformation describe-stacks \
        --stack-name asset-price-service \
        --query 'Stacks[0].Outputs[?OutputKey==`AssetPriceApiUrl`].OutputValue' \
        --output text)
    
    echo "🌐 API Endpoint: $API_URL"
    
else
    echo ""
    echo "❌ Deployment failed"
    exit 1
fi
