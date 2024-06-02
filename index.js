import grpc from "k6/net/grpc";
import { config } from "./k6/entity/config.js";
import { initMerchantNearestLocations, initPegeneratedTSPMerchants, initZonesWithPregeneratedMerchants, resetAll } from "./k6/repository/initCaddyRepository.js";
import { assignPregeneratedMerchant, getAllMerchantNearestLocations, getAllMerchantRoutes, getAllPregeneratedMerchants } from "./k6/repository/getterCaddyRepository.js";
import { AdminRegisterTest } from "./k6/tests/auth/adminRegisterTest.js";
import { fail } from "k6";
import { AdminLoginTest } from "./k6/tests/auth/adminLoginTest.js";
import { UserRegisterTest } from "./k6/tests/auth/userRegisterTest.js";
import { UserLoginTest } from "./k6/tests/auth/userLoginTest.js";
import { MerchantPostTest } from "./k6/tests/merchantManagement/merchantPostTest.js";
import { MerchantGetTest } from "./k6/tests/merchantManagement/merchantGetTest.js";
import { MerchantItemGetTest } from "./k6/tests/merchantManagement/merchantItemGetTest.js";
import { MerchantItemPostTest } from "./k6/tests/merchantManagement/merchantItemPostTest.js";
import { EstimateOrderTest } from "./k6/tests/purchase/estimateOrderTest.js";
import { GetNearbyMerchantTest } from "./k6/tests/purchase/getNearbyMerchantTest.js";
import { OrderTest } from "./k6/tests/purchase/orderTest.js";
import { GetOrderTest } from "./k6/tests/purchase/getOrderTest.js";

const client = new grpc.Client();
client.load([], 'contract.proto');

export function setup() {
    client.connect('127.0.0.1:50051', {
        plaintext: true
    });

    initZonesWithPregeneratedMerchants(client);
    initMerchantNearestLocations(client);
    initPegeneratedTSPMerchants(client);

    client.close();
}
export function teardown() {
    client.connect('127.0.0.1:50051', {
        plaintext: true
    });

    resetAll(client);

    client.close();
}

export default function () {
    client.connect('127.0.0.1:50051', {
        plaintext: true
    });

    if (config.LOAD_TEST) {
        return // stub
    }

    // Admin test
    let admin = AdminRegisterTest(config, { feature: "Admin Register" })
    checkRes(admin, "Admin Register failed")
    admin = AdminLoginTest(admin, config, { feature: "Admin Login" })
    checkRes(admin, "Admin Login failed")

    const merchantToAdd = getAllPregeneratedMerchants(client)
    let merchants = []

    let merchant = MerchantPostTest(admin, merchantToAdd.merchant.pop(), config, { feature: "Merchant Post" })
    checkRes(merchant, "Merchant Post failed")
    merchants.push(merchant)

    merchantToAdd.merchant = merchantToAdd.merchant.slice(0, merchantToAdd.merchant.length - 1)
    merchant = MerchantGetTest(admin, merchantToAdd.merchant, config, { feature: "Merchant Get" })
    checkRes(merchant, "Merchant Get failed")
    merchants.push(merchant)

    merchants.forEach(m => {
        assignPregeneratedMerchant(client, { merchantId: m.merchantId, pregeneratedId: m.pregeneratedId })
    });

    const merchantToAddItem = merchants.shift()
    let merchantItems = MerchantItemGetTest(admin, merchantToAddItem, 10, config, { feature: "Merchant Item Get" })
    checkRes(merchantItems, "Merchant Item Get failed")

    let merchantItem = MerchantItemPostTest(admin, merchantToAddItem, config, { feature: "Merchant Item Post" })
    checkRes(merchantItem, "Merchant Item Post failed")
    merchantItems.push(merchantItem)


    // User test
    let user = UserRegisterTest(config, { feature: "User Register" })
    checkRes(user, "User Register failed")
    user = UserLoginTest(user, config, { feature: "User Login" })
    checkRes(user, "User Login failed")

    let allNearestMerchantLocations = getAllMerchantNearestLocations(config)
    checkRes(allNearestMerchantLocations, "Get all nearest merchant locations failed")
    GetNearbyMerchantTest(user, allNearestMerchantLocations.records[0], config, { feature: "Get Nearby Merchant" })

    let allMerchantRoutes = getAllMerchantRoutes(config)
    checkRes(allMerchantRoutes, "Get all merchant routes failed")
    const estimateId = EstimateOrderTest(user, admin, allMerchantRoutes.zone[0], allMerchantRoutes.zone[1], config, { feature: "Estimate Order" })
    checkRes(estimateId, "Estimate Order failed")
    OrderTest(user, estimateId, config, { feature: "Order" })
    GetOrderTest(user, config, { feature: "Get Order" })

    client.close()
}

function checkRes(res, errMsg) {
    if (!res) {
        fail(errMsg)
    }
}