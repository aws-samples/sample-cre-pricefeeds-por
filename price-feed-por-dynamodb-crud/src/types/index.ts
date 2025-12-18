export interface AssetPriceRecord {
    recordType: 'ASSET_PRICE';
    timestamp: string;
    price: number;
}

export interface ProofOfReservesRecord {
    recordType: 'PROOF_OF_RESERVES';
    timestamp: string;
    collateralUsd: number;
}

export type DataRecord = AssetPriceRecord | ProofOfReservesRecord;
