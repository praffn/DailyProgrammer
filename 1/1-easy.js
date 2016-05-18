const readline = require('readline');
const rl = readline.createInterface({
  input: process.stdin,
  output: process.stdout
});

rl.question('Enter your name\n', name => {
  rl.question('Enter your age\n', age => {
    rl.question('Enter your username\n', username => {
      console.log(`Your name is ${name}, you are ${age} and your username is ${username}`);
      rl.close();
    });
  });
});
