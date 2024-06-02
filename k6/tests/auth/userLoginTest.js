import { IsUser } from "../../entity/user.js";
import { combine, generateTestObjects } from "../../helpers/generator.js";
import { testPostJsonAssert } from "../../helpers/request.js";

/**
 * @param {Admin} user
 * @param {import("../../entity/config").Config} config 
 * @returns {null | import("../../entity/user").User}
 */
export function UserLoginTest(user, config, tags) {
    if (!IsUser(user)) {
        return;
    }

    const featureName = "User Login";
    const route = config.BASE_URL + "/user/login";

    /** @type {import("../../entity/user").User} */
    const positivePayload = {
        username: user.username,
        password: user.password,
    }

    if (!config.POSITIVE_CASE) {
        testPostJsonAssert(
            "empty body",
            featureName,
            route, {}, {}, {
            ['should return 400']: (v) => v.status === 400
        }, config, tags)

        const testObjects = generateTestObjects({
            username: { type: "string", notNull: true, minLength: 5, maxLength: 30 },
            password: { type: "string", notNull: true, minLength: 5, maxLength: 30 },
        }, positivePayload)
        testObjects.forEach(payload => {
            testPostJsonAssert("invalid payload", featureName, route, payload, {}, {
                ['should return 400']: (res) => res.status === 400,
            }, config, tags);
        });
    }

    const res = testPostJsonAssert("valid payload", featureName, route, positivePayload, {}, {
        ['should return 200']: (v) => v.status === 200
    }, config, tags)

    if (res.isSuccess) {
        return combine(user, {
            token: res.res.json().token
        })
    }
}