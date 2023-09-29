const upperCase = require("upper-case").upperCase;

function greet(name) {
  console.log(upperCase(`Hello ${name}, welcome to Tester`));
}

module.exports = greet;
