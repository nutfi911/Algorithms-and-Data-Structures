// Both numbers have the same digits (order doesn't matter)

// Return: true/false

// ### Examples

// 182, 281 -> True
// 34, 41 -> False
// 3589578, 5879385 -> True
// 22, 222 -> False

// TODO: Convert to GO
function sameFrequency(a, b) {
    if (a.length !== b.length) {
        return [a, b, false];
    }

    const aFrequencies = {};
    const bFreqnuencies = {};

    for (let i = 0; i < a.length; i++) {
        if (aFrequencies[a[i]]) {
            aFrequencies[a[i]]++;
        } else {
            aFrequencies[a[i]] = 1;
        }
        if (bFreqnuencies[b[i]]) {
            bFreqnuencies[b[i]]++;
        } else {
            bFreqnuencies[b[i]] = 1;
        }
    }

    for (const key in aFrequencies) {
        if (aFrequencies[key] !== bFreqnuencies[key]) {
            return [a, b, false];
        }
    }
    return [a, b, true];
}
console.log(...sameFrequency([1, 8, 2], [2, 8, 1]));
console.log(...sameFrequency([3, 4], [4, 1]));
console.log(...sameFrequency([3, 5, 8, 9, 5, 7, 8], [5, 8, 7, 9, 3, 8, 5]));
console.log(...sameFrequency([2, 2], [2, 2, 2]));
