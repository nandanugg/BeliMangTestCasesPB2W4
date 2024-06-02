import grpc from "k6/net/grpc";

/**
 * 
 * @param {grpc.Client} grpcClient
 * @returns {null | import("../entity/merchant").MerchantNearestRecordZone}
 */
export function getAllMerchantNearestLocations(grpcClient) {
    const resp = grpcClient.invoke("pb.MerchantService/GetAllMerchantNearestLocations", {});
    if (resp.status !== grpc.StatusOK) {
        logGrpcError(resp)
        return null
    }
    return resp.message
}

/**
 * 
 * @param {grpc.Client} grpcClient 
 * @returns {null | import("../entity/merchant").AllGeneratedRoutes}
 */
export function getAllMerchantRoutes(grpcClient) {
    const resp = grpcClient.invoke("pb.MerchantService/GetAllMerchantRoutes", {});
    if (resp.status !== grpc.StatusOK) {
        logGrpcError(resp)
        return null
    }
    return resp.message
}

/**
 * 
 * @param {grpc.Client} grpcClient 
 * @returns {null | import("../entity/merchant").PregeneratedMerchant}
 */
export function getAllPregeneratedMerchants(grpcClient) {
    const resp = grpcClient.invoke("pb.MerchantService/GetAllPregeneratedMerchants", {});
    if (resp.status !== grpc.StatusOK) {
        logGrpcError(resp)
        return null
    }
    return resp.message
}

/**
 * 
 * @param {grpc.Client} grpcClient 
 * @param {import("../entity/merchant").AssignMerchant} assignMerchant
 * @returns {null | Object}
 */
export function assignPregeneratedMerchant(grpcClient, assignMerchant) {
    const resp = grpcClient.invoke("pb.MerchantService/AssignPregeneratedMerchant", assignMerchant);
    if (resp.status !== grpc.StatusOK) {
        logGrpcError(resp)
        return null
    }
    return resp.message
}


/**
 * 
 * @param {import("k6/net/grpc").Response} resp 
 */
export function logGrpcError(resp) {
    console.error(`Error grpc: code = ${resp.status}, msg = ${resp.message}`);
}