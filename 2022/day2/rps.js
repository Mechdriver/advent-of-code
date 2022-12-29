const fs = require("fs");

const data = fs.readFileSync("./input.txt", "utf8");

const moves = data.split("\n");

const OPPONENT = {
  rock: "A",
  paper: "B",
  scissors: "C",
};

const PLAYER = {
  rock: "X",
  paper: "Y",
  scissors: "Z",
};

const SCORES = {
  A: 1,
  B: 2,
  C: 3,
};

const SCORES_P2 = {
  A: { X: "C", Y: "A", Z: "B" },
  B: { X: "A", Y: "B", Z: "C" },
  C: { X: "B", Y: "C", Z: "A" },
  X: 0,
  Y: 3,
  Z: 6,
};

let score = 0;

for (let game of moves) {
  const opponent = game[0];
  const result = game[2];
  score += SCORES_P2[result];
  const move =  SCORES_P2[opponent][result];
  score += SCORES[move];
}

console.log(score);
