import * as fs from "fs";

const data = fs.readFileSync("./input.txt", "utf8");

const rawReportRows = data.split("\n").filter((row) => !!row);

const reportRows = rawReportRows.map((row) => {
  return row
    .split(/[ ]+/)
    .flat()
    .map((val) => Number(val));
});

let safeSum = 0;

const isSafe = (row) => {
  let isIncreasing;

  for (let i = 0, j = 1; i < row.length; i++, j++) {
    const delta = Math.abs(row[i] - row[j]);
    if (delta < 1 || delta > 3) {
      return false;
    }

    if (isIncreasing === undefined) {
      if (row[i] < row[j]) {
        isIncreasing = true;
      } else {
        isIncreasing = false;
      }
    } else {
      if (isIncreasing) {
        if (row[i] > row[j]) {
          return false;
        }
      } else {
        if (row[i] < row[j]) {
          return false;
        }
      }
    }
  }

  return true;
};

reportRows.map((row) => {
  if (isSafe(row)) {
    safeSum += 1;
  }
});

console.log(safeSum);
