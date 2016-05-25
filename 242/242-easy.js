const input = process.argv.slice(2);
const people = parseInt(input[0], 10);
const plants = parseInt(input[1], 10);

function plant (people, plants) {
  'use strict';
  let weeks = 0;
  let fruits = 0;
  while (fruits < people) {
    weeks++;
    fruits += plants;
    plants += fruits;
  }
  return weeks + 1;
}

console.log(plant(people, plants));
