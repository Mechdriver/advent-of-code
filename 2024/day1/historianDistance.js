import * as fs from "fs";

const data = fs.readFileSync("./input.txt", "utf8");

const locationIds = data.split("\n");

const comparator = (a, b) => a - b;
const leftList = [];
const rightList = [];
let distance = 0;

locationIds.map((row) => {
  if (row) {
    const splitIds = row.split(/[ ]+/);

    leftList.push(splitIds[0]);
    rightList.push(splitIds[1]);
  }
});

leftList.sort(comparator);
rightList.sort(comparator);

for (let i = 0; i < leftList.length; i++) {
  distance += Math.abs(leftList[i] - rightList[i]);
}

console.log(distance);
