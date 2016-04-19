// simple arithmetic calculator
// use like so:
// node 2.js <X> <OPERATOR> <Y>
// Legal operators:
// + : addition
// - : subtraction
// x : multiplication (*/asterisk cannot be used unquoted, and thus I settled with x)
// / : division

var x = parseInt(process.argv[2], 10);
var operator = process.argv[3];
var y = parseInt(process.argv[4], 10);

function add (x, y) {
  return x + y;
}

function subtract (x, y) {
  return x - y;
}

function multiply (x, y) {
  return x * y;
}

function divide (x, y) {
  return x / y;
}

var operations = {
  '+': add,
  '-': subtract,
  'x': multiply,
  '/': divide
};

console.log(operations[operator](x, y));
