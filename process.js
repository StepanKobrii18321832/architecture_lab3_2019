const fs = require('fs');

const path_from = process.argv[2];
const path_to = process.argv[3];

fs.exists(path_to, exists => {
  if (!exists)
    fs.mkdir(path_to, err => {
      if (err) throw err
    });
});

fs.readdir(path_from, (err, files) => {
  files.map(file => {
    fs.readFile(path_from + file, 'utf8', (err, data) => {
      fs.writeFile(
        path_to + file.split('.')[0] + '.res',
        data.split(/(?<!\w\.\w.)(?<![A-Z][a-z]\.)(?<=\.|\?)\s/).length,
        () => { }
      );
    });
  });
  console.log(`Total number of processed files: ${files.length}`);
});

