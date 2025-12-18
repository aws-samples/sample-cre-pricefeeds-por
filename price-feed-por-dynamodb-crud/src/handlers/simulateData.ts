import { APIGatewayProxyEvent, APIGatewayProxyResult } from 'aws-lambda';
import { DynamoDBClient } from '@aws-sdk/client-dynamodb';
import { DynamoDBDocumentClient, BatchWriteCommand } from '@aws-sdk/lib-dynamodb';
import { AssetPriceRecord, ProofOfReservesRecord } from '../types';

const client = new DynamoDBClient({});
const docClient = DynamoDBDocumentClient.from(client);
const TABLE_NAME = process.env.TABLE_NAME || 'AssetData';

// Base price for the simulated asset
const BASE_PRICE = 1000;

// Collateral ratio range (120% to 150%)
const MIN_COLLATERAL_RATIO = 1.2;
const MAX_COLLATERAL_RATIO = 1.5;

// Price variation range (-5% to +5%)
const PRICE_VARIATION = 0.05;

/**
 * Lambda handler for generating and storing simulated asset price and proof of reserves data
 * Requirements: 1.1, 3.1 (supports data generation for testing and demonstration)
 */
export async function handler(event: APIGatewayProxyEvent): Promise<APIGatewayProxyResult> {
    try {
        // Generate simulated asset price using random walk algorithm
        const randomVariation = (Math.random() * 2 - 1) * PRICE_VARIATION; // Random value between -0.05 and +0.05
        const simulatedPrice = BASE_PRICE * (1 + randomVariation);
        const price = Math.round(simulatedPrice * 100) / 100; // Round to 2 decimal places

        // Generate collateral amount with 120-150% ratio to price
        const collateralRatio = MIN_COLLATERAL_RATIO + Math.random() * (MAX_COLLATERAL_RATIO - MIN_COLLATERAL_RATIO);
        const collateralUsd = Math.round(price * collateralRatio * 100) / 100; // Round to 2 decimal places

        // Create timestamp
        const timestamp = new Date().toISOString();

        // Create asset price record
        const assetPriceRecord: AssetPriceRecord = {
            recordType: 'ASSET_PRICE',
            timestamp,
            price
        };

        // Create proof of reserves record
        const proofOfReservesRecord: ProofOfReservesRecord = {
            recordType: 'PROOF_OF_RESERVES',
            timestamp,
            collateralUsd
        };

        // Store both records atomically using BatchWriteItem
        await docClient.send(new BatchWriteCommand({
            RequestItems: {
                [TABLE_NAME]: [
                    {
                        PutRequest: {
                            Item: assetPriceRecord
                        }
                    },
                    {
                        PutRequest: {
                            Item: proofOfReservesRecord
                        }
                    }
                ]
            }
        }));

        // Return success response with generated data
        return {
            statusCode: 200,
            headers: {
                'Content-Type': 'application/json',
                'Access-Control-Allow-Origin': '*'
            },
            body: JSON.stringify({
                price,
                collateralUsd,
                timestamp
            })
        };

    } catch (error) {
        console.error('Error simulating data:', error);

        return {
            statusCode: 500,
            headers: {
                'Content-Type': 'application/json',
                'Access-Control-Allow-Origin': '*'
            },
            body: JSON.stringify({
                error: 'Internal server error',
                code: 'INTERNAL_ERROR'
            })
        };
    }
}
