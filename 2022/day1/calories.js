const fs = require('fs');

const data = fs.readFileSync('./input.txt', 'utf8');

let calList = data.split('\n');
calList = calList.map((item) => {
  if (item !== '') {
    return Number(item);
  }
  return null;
});


let sumCals = [];
let sum = 0;

for (let snack of calList) {
  if (snack === null) {
    sumCals.push(sum);
    sum = 0;
  } else {
    sum += snack;
  }
}

sumCals.sort();
sumCals = sumCals.slice(-4, -1);


const maxCals = sumCals.reduce((a, b) => a + b);
console.log(maxCals);
