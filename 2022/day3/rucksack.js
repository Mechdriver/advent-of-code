const fs = require("fs");
const data = fs.readFileSync("./input.txt", "utf8");
let rucksacks = data.split("\n");

const lowerBase = 96;
const upperBase = 64;
const priorityMap = new Map();
const priorityList = [];

function getSackPriorities(rucksacks) {
  for (let sack of rucksacks) {
    const itemSet = new Set();
    const leftHalf = sack[0];
    const rightHalf = sack[1];

    for (let item of leftHalf) {
      itemSet.add(item);
    }

    for (let item of rightHalf) {
      if (itemSet.has(item)) {
        priorityList.push(priorityMap.get(item));
        break;
      }
    }
  }
  return priorityList;
}

for (let i = lowerBase + 1; i < 123; i++) {
  priorityMap.set(String.fromCharCode(i), i - lowerBase);
}

for (let i = upperBase + 1; i < 91; i++) {
  priorityMap.set(String.fromCharCode(i), i - upperBase + 26);
}

// rucksacks = rucksacks.map((sack) => {
//   return [sack.slice(0, sack.length / 2), sack.slice(sack.length / 2)];
// });

for (let i = 0; i < rucksacks.length; i += 3) {
  const sack1 = rucksacks[i];
  const sack2 = rucksacks[i + 1];
  const sack3 = rucksacks[i + 2];
  const itemSet = new Set();
  const itemSet2 = new Set();

  for (let item of sack1) {
    itemSet.add(item);
  }

  for (let item of sack2) {
    if (itemSet.has(item)) {
      itemSet2.add(item);
    }
  }

  for (let item of sack3) {
    if (itemSet2.has(item)) {
      priorityList.push(priorityMap.get(item));
      break;
    }
  }
}

console.log(priorityList.reduce((a, b) => a + b));
