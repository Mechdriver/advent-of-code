import * as fs from "fs";

const data = fs.readFileSync("./input.txt", "utf8");

const locationIds = data.split("\n");

const leftList = [];
const rightMap = new Map();
let similarity = 0;

locationIds.map((row) => {
  if (row) {
    const splitIds = row.split(/[ ]+/);
    const rightId = splitIds[1];

    leftList.push(splitIds[0]);

    if (rightMap.has(rightId)) {
      rightMap.set(rightId, rightMap.get(rightId) + 1);
    } else {
      rightMap.set(rightId, 1);
    }
  }
});

for (let i = 0; i < leftList.length; i++) {
  const locationId = leftList[i];
  const multiplier = rightMap.has(locationId) ? rightMap.get(locationId) : 0;

  similarity += locationId * multiplier;
}

console.log(similarity);
