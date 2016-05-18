const fs = require('fs');
const path = require('path');

// read the input file (sync)
var input = fs.readFileSync(path.join(__dirname, '266-easy-input.txt'), 'utf8');
input = input
  .split('\n')
  .map(n => n.split(' ').map(i => Number(i)))

// get node count
var nodeCount = input[0][0];
// and create a zerofilled array with an element for each node
var nodes = new Array(nodeCount).fill(0);

// eww... Mutate the fuck out of that input, and remove the first element.
input.splice(0, 1);
// and flatten it out
var flattened = [];
input.forEach(a => {
  flattened = flattened.concat(a);
});

// and now we just count :)
flattened.forEach(n => {
  nodes[n - 1]++
});

// and finally we output the results
nodes.forEach((n, i) => {
  console.log(`Node ${i + 1} has a degree of ${n}`);
});
