const fs = require('fs');

const path_from = process.argv[2];
const path_to = process.argv[3];

fs.exists(path_to, exists => {
    if (!exists) fs.mkdir(path_to, err => {if (err) throw err});
});

fs.readdir(path_from, function(err, items) {
    console.log(items);
    items.map(i => {
        fs.readFile(path_from + i, 'utf8', function(err, data = []) {
            let count = 0;
            //console.log(i);
            for (let j = 1; j < data.length; j++) {
                if (((data[j] === '.') || (data[j] === '!') || (data[j] === '?') ||
                    ((data[j] === '.') && (data[j - 1] === '.') && data[j - 2] === '.')) &&
                    ((data[j + 1] === ' ') || ((data[j + 1]) === '\n') || (!data[j + 1]))) count++;
                fs.writeFile(path_to + i.split('.')[0] + '.res', count, () => {});
            }
        });
    })
    console.log("Total number of processed files: " + items.length);
});
