const fs = require('fs');
const buffer = fs.readFileSync('input.txt');
const input = buffer.toString().split('\n');

const numbers = [
  'one',
  'two',
  'three',
  'four',
  'five',
  'six',
  'seven',
  'eight',
  'nine',
];

const result = (input) => {
  let ans = 0;

  input.forEach((word) => {
    const value = getCalibrationValue(word);
    ans += value;
  });

  return ans;
};

const getCalibrationValue = (word) => {
  let firstNumber = '';
  let lastNumber = '';
  let numberWord = '';

  word.split('').forEach((letter) => {
    if (!isNaN(letter)) {
      if (!firstNumber) {
        firstNumber = letter;
        lastNumber = letter;
      } else {
        lastNumber = letter;
      }
    } else {
      numberWord += letter;
      const matches = numbers.filter((number) => {
        return number.startsWith(numberWord);
      });

      if (matches.length) {
        if (numberWord === matches[0]) {
          const number = numbers.indexOf(numberWord) + 1;
          if (!firstNumber) {
            firstNumber = number.toString();
            lastNumber = number.toString();
          } else {
            lastNumber = number.toString();
          }

          numberWord = letter;
        }
      } else {
        numberWord = numberWord.slice(1);
      }
    }
  });

  return Number(firstNumber + lastNumber);
};

console.log(result(input));
