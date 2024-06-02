import grpc from "k6/net/grpc";
import { logGrpcError } from "./getterCaddyRepository.js";

/**
 * 
 * @param {grpc.Client} grpcClient 
 * @returns {null | Object}
 */
export function initMerchantNearestLocations(grpcClient) {
    const resp = grpcClient.invoke("pb.MerchantService/InitMerchantNearestLocations", {});
    if (resp.status !== grpc.StatusOK) {
        logGrpcError(resp)
        return null
    }
    return resp.message
}

/**
 * 
 * @param {grpc.Client} grpcClient 
 * @returns {null | Object}
 */
export function initPegeneratedTSPMerchants(grpcClient) {
    const resp = grpcClient.invoke("pb.MerchantService/InitPegeneratedTSPMerchants", {});
    if (resp.status !== grpc.StatusOK) {
        logGrpcError(resp)
        return null
    }
    return resp.message
}

/**
 * 
 * @param {grpc.Client} grpcClient 
 * @returns {null | Object}
 */
export function initZonesWithPregeneratedMerchants(grpcClient) {
    const resp = grpcClient.invoke("pb.MerchantService/InitZonesWithPregeneratedMerchants", {});
    if (resp.status !== grpc.StatusOK) {
        logGrpcError(resp)
        return null
    }
    return resp.message
}

/**
 * 
 * @param {grpc.Client} grpcClient 
 * @returns {null | Object}
 */
export function resetAll(grpcClient) {
    const resp = grpcClient.invoke("pb.MerchantService/ResetAll", {});
    if (resp.status !== grpc.StatusOK) {
        logGrpcError(resp)
        return null
    }
    return resp.message
}