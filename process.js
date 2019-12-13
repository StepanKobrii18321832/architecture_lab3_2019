const fs = require('fs');

const path_from = process.argv[2];
const path_to = process.argv[3];

fs.exists(path_to, exists => {
    if (!exists) fs.mkdir(path_to, err => {if (err) throw err});
});

fs.readdir(path_from, function(err, items) {
    console.log(items);
    

    
    console.log("Total number of processed files: " + items.length);
});

