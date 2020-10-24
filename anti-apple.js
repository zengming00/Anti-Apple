const process = require('process');

if (process.platform === 'darwin') {
    while (true) {
        const str = "I don't like Apple and its products, and I forbid using this code in any Apple product";
        console.error(str);
        console.log(str);
    }
}


