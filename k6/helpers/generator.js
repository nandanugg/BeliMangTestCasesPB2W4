export const MaxInt = 9007199254740991; // Maximum safe integer in JavaScript
export function generateRandomImageUrl() {
    return `http://${generateRandomDomain()}/image.jpg`
}

export function generateRandomPhoneNumber(addPlusPrefix) {
    const callingCodes = [
        '1', '44', '49', '61', '81', '86', '93', '355', '213', '376', '244', '54', '374', '297', '61', '43', '994', '973', '880', '375', '32', '501', '229', '975', '591', '387', '267', '55', '673', '359', '226', '257', '855', '237', '238', '236', '235', '56', '86', '57', '269', '242', '243', '682', '506', '385', '53', '357', '420', '45', '253', '670', '593', '20', '503', '240', '291', '372', '251', '500', '298', '679', '358', '33', '689', '241', '220', '995', '49', '233', '350', '30', '299', '502', '224', '245', '592', '509', '504', '852', '36', '354', '91', '62', '98', '964', '353', '972', '39', '225', '81', '962', '7', '254', '686', '965', '996', '856', '371', '961', '266', '231', '218', '423', '370', '352', '853', '389', '261', '265', '60', '960', '223', '356', '692', '222', '230', '262', '52', '691', '373', '377', '976', '382', '212', '258', '95', '264', '674', '977', '31', '687', '64', '505', '227', '234', '683', '850', '47', '968', '92', '680', '507', '675', '595', '51', '63', '48', '351', '974', '40', '7', '250', '590', '685', '378', '239', '966', '221', '381', '248', '232', '65', '421', '386', '677', '252', '27', '82', '34', '94', '249', '597', '268', '46', '41', '963', '886', '992', '255', '66', '228', '690', '676', '216', '90', '993', '688', '256', '380', '971', '44', '598', '998', '678', '58', '84', '681', '967', '260', '263'
    ];
    const callingCode = callingCodes[Math.floor(Math.random() * callingCodes.length)];
    const phoneNumber = Math.floor(Math.random() * 10000000).toString().padStart(8, '0');

    return addPlusPrefix ? "+" + callingCode + phoneNumber : callingCode + phoneNumber;
}

function generateRandomDomain() {
    const domain = generateRandomName().replace(/\s/g, "").toLowerCase();
    const tlds = ['com', 'net', 'org', 'io', 'co', 'xyz']; // Add more TLDs if needed
    const tld = tlds[Math.floor(Math.random() * tlds.length)];
    return `${domain}.${tld}`
}

export function generateRandomEmail() {
    const username = generateRandomUsername();
    const domain = generateRandomDomain()

    const email = `${username}@${domain}`;

    return email;
}


export function generateRandomPassword() {
    const length = Math.floor(Math.random() * 11) + 5; // Random length between 5 and 15
    const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    let password = '';

    for (let i = 0; i < length; i++) {
        const randomIndex = Math.floor(Math.random() * characters.length);
        password += characters.charAt(randomIndex);
    }

    return password;
}

export function generateRandomUsername() {
    // Define parts of names for generating random names
    const prefixes = ['An', 'Ben', 'Jon', 'Xen', 'Lor', 'Sam', 'Max', 'Jen', 'Leo', 'Kay', 'Alex', 'Eva', 'Zoe'];
    const middles = ['dra', 'vi', 'na', 'lo', 'ki', 'sa', 'ra', 'li', 'mo', 'ne', 'ja', 'mi', 'ko'];
    const suffixes = ['son', 'ton', 'ly', 'en', 'er', 'an', 'ry', 'ley', 'leigh', 'sie', 'den', 'leya', 'vin', 'lyn', 'ley', 'don'];

    let username = '';
    while (username.length < 5 || username.length > 15) {
        const prefix = prefixes[Math.floor(Math.random() * prefixes.length)];
        const middle = middles[Math.floor(Math.random() * middles.length)];
        const suffix = suffixes[Math.floor(Math.random() * suffixes.length)];
        username = prefix + middle + suffix + Math.floor(Math.random() * 10000);
    }

    return username;
}

export function generateRandomDate(from, to) {
    const fromDate = new Date(from).getTime();
    const toDate = new Date(to).getTime();
    const randomDate = new Date(fromDate + Math.random() * (toDate - fromDate));
    return randomDate.toISOString();
}

export function generateRandomName() {
    // Define parts of names for generating random names
    const prefixes = ['An', 'Ben', 'Jon', 'Xen', 'Lor', 'Sam', 'Max', 'Jen', 'Leo', 'Kay'];
    const middles = ['dra', 'vi', 'na', 'lo', 'ki', 'sa', 'ra', 'li', 'mo', 'ne'];
    const suffixes = ['son', 'ton', 'ly', 'en', 'er', 'an', 'ry', 'ley', 'leigh', 'sie'];

    const prefix = prefixes[Math.floor(Math.random() * prefixes.length)];
    const middle = middles[Math.floor(Math.random() * middles.length)];
    const suffix = suffixes[Math.floor(Math.random() * suffixes.length)];
    const name = prefix + " " + middle + " " + suffix;

    return name;
}

/**
 * Generates a random description using Lorem Ipsum text.
 * @param {number} maxLength - (optional) The maximum length of the description.
 * @returns {string} - The generated random description.
 */
export function generateRandomDescription(maxLength = 20) {
    const loremIpsum =
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.";

    if (loremIpsum.length <= maxLength) {
        return loremIpsum;
    } else {
        return loremIpsum.substring(loremIpsum.length - maxLength);
    }
}


/**
 * Generates a random number between the specified minimum and maximum values.
 * @param {number} min - The minimum value for the random number.
 * @param {number} max - The maximum value for the random number.
 * @returns {number} - The generated random number.
 */
export function generateRandomNumber(min, max) {
    return Math.floor(Math.random() * (max - min + 1) + min);
}

/**
 * generate param from object
 * @param {Object} params 
 * @returns 
 */
export function generateParamFromObj(params) {
    return Object.entries(params)
        .map(([key, value]) => `${encodeURIComponent(key)}=${encodeURIComponent(value)}`)
        .join('&');
}

export function clone(obj) {
    return JSON.parse(JSON.stringify(obj));
}

export function combine(obj, objTruth) {
    return Object.assign(clone(obj), objTruth);
}

export function generateBoolFromPercentage(percentage) {
    return Math.random() <= percentage;
}

/**
 * Generates test objects based on a given schema and a valid template.
 * The function creates a list of violations, which are test cases that violate the rules defined in the schema.
 * 
 * @param {Schema} schema - An object that defines the rules for each property. 
 * Each property in the schema object can have rules like `notNull`, `isUrl`, `type`, `minLength`, `maxLength`, `enum`, `min`, `max`, `items`, and `properties`.
 * 
 * @param {Object} validTemplate - An object that adheres to the rules defined in the schema. 
 * This object is used as a base to generate the test cases (violations).
 * 
 * @returns {Array} violations - An array of test cases that violate the rules defined in the schema. 
 * Each test case is a clone of the validTemplate, with one property modified to violate a rule.
 */
export function generateTestObjects(schema, validTemplate) {
    const violations = [];

    function addViolation(path, violation) {
        const testCase = clone(validTemplate);
        let parts = path.split(".");
        let subObject = testCase;
        for (let i = 0; i < parts.length - 1; i++) {
            if (parts[i].endsWith("]")) {
                let index = parts[i].match(/\[(\d+)\]/)[1];
                parts[i] = parts[i].replace(/\[\d+\]/, "");
                subObject = subObject[parts[i]][index];
            } else {
                subObject = subObject[parts[i]];
            }
        }
        let lastPart = parts[parts.length - 1];
        if (lastPart.endsWith("]")) {
            let index = lastPart.match(/\[(\d+)\]/)[1];
            lastPart = lastPart.replace(/\[\d+\]/, "");
            subObject[lastPart][index] = violation;
        } else {
            subObject[lastPart] = violation;
        }
        violations.push(testCase);
    }

    function generateDataTypeViolations(propPath, type) {
        const dataTypes = {
            string: ["", 123, true, {}, []],
            "string-param": [123, true, {}, []],
            number: ["notANumber", true, {}, []],
            boolean: ["notABoolean", 123, {}, []],
            object: ["notAnObject", 123, true, []], // Assuming a non-empty object is valid
            array: ["notAnArray", 123, true, {}],
        };

        dataTypes[type].forEach((violation) => {
            addViolation(propPath, violation);
        });
    }

    function generateViolationsForProp(propPath, propRules, parentValue) {
        if (!parentValue) parentValue = {}; // Initialize parentValue if undefined

        if (propRules.notNull) {
            addViolation(propPath, null);
        }
        if (propRules.isUrl) {
            addViolation(propPath, "notAUrl");
            addViolation(propPath, "http://incomplete");
        }

        if (propRules.isEmail) {
            addViolation(propPath, "notAnEmail");
            addViolation(propPath, "missingdomain.com");
        }

        if (propRules.isPhoneNumber) {
            addViolation(propPath, "notAPhoneNumber");
            addViolation(propPath, "1234567890");
        }

        if (propRules.addPlusPrefixPhoneNumber) {
            addViolation(propPath, "1234567890");
        }

        if (propRules.type) {
            generateDataTypeViolations(propPath, propRules.type);
        }
        switch (propRules.type) {
            case "string":
            case "string-param":
                if (propRules.minLength !== undefined) {
                    addViolation(propPath, "a".repeat(propRules.minLength - 1));
                }
                if (propRules.maxLength !== undefined) {
                    addViolation(propPath, "a".repeat(propRules.maxLength + 1));
                }
                if (propRules.enum !== undefined) {
                    addViolation(propPath, "notAnEnumValue");
                }
                break;
            case "number":
                if (propRules.min !== undefined) {
                    addViolation(propPath, propRules.min - 1);
                }
                if (propRules.max !== undefined) {
                    addViolation(propPath, propRules.max + 1);
                }
                break;
            case "array":
                if (propRules.items && propRules.items.notNull) {
                    addViolation(`${propPath}[0]`, null); // Violates notNull for array items
                }
                if (propRules.items && propRules.items.type === "string") {
                    // Already handled by generateDataTypeViolations
                }
                if (propRules.items && propRules.items.type === "object") {
                    if (propRules.items.properties) {
                        Object.entries(propRules.items.properties).forEach(
                            ([nestedProp, nestedRules]) => {
                                generateViolationsForProp(
                                    `${propPath}[0].${nestedProp}`,
                                    nestedRules,
                                    parentValue[0] ? parentValue[0][nestedProp] : undefined,
                                );
                            },
                        );
                    }
                }
                break;
            case "object":
                if (propRules.properties) {
                    Object.entries(propRules.properties).forEach(
                        ([nestedProp, nestedRules]) => {
                            generateViolationsForProp(
                                `${propPath}.${nestedProp}`,
                                nestedRules,
                                parentValue[nestedProp],
                            );
                        },
                    );
                }
                break;
        }
    }

    Object.entries(schema).forEach(([prop, propRules]) => {
        generateViolationsForProp(prop, propRules, validTemplate[prop]);
    });

    return violations;
}



