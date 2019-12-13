'use strict';
const fs = require('fs');
const inputDir = process.argv[2];
const outputDir = process.argv[3];

const lab3 = (input, output) => {
  if (!fs.existsSync(input)) return null;
  if (!fs.existsSync(output)) fs.mkdirSync(output);

  fs.readdir(input, 'utf-8', (err, files) => {
    if (err) throw err;
    Promise.all(files.map(file => {
      return new Promise((resolve, reject) => {
        let count = 0;
        let result = '';
        const readstream = fs.createReadStream(`./${input}/${file}`, err => { if (err) reject(err); });
        readstream.on('data', data => {
          result += data;
          count = result.split(/(?<!\w\.\w.)(?<![A-Z][a-z]\.)(?<=\.|\?)\s/).length;
        })
        readstream.on('end', () => {
          fs.writeFile(
            `${output}/${file.split('.')[0]}.res`,
            count,
            'utf-8',
            err => { if (err) console.log(err); });
        });
        resolve(true);
      });
    }))
      .then(res => {
        console.log(`Total number of processed files: ${res.length}.`);
      }, err => { if (err) console.log(err.stack); });
  });
};

lab3(inputDir, outputDir);
