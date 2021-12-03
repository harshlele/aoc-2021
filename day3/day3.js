const fs = require("fs");
const txt = fs.readFileSync("day3/main-input.txt");

const lines = txt.toString().split("\r\n");

let gamma = [], epsilon = [];
let ones = new Array(lines[0].length).fill(0);
//part 1
for(let i = 0; i < lines.length; i++){
    for(let j = 0; j < lines[i].length; j++){
        if(lines[i][j] == "1"){
            ones[j] += 1;
        }   
        
        if(ones[j] > lines.length - ones[j]) {
            gamma[j] = 1;
            epsilon[j] = 0;
        }
        else {
            gamma[j] = 0;
            epsilon[j] = 1;
        }
    }
}
let g = parseInt(gamma.join(""),2);
let e = parseInt(epsilon.join(""),2);
console.log(g * e);


let oFilter = [], cFilter = [];

lines.forEach((l,i) => {
    if(gamma[0] == parseInt(l[0])) oFilter.push(i);
    else cFilter.push(i);
});

let oDone = false,cDone = false;
for(i = 1; i < gamma.length; i++){
    let ones = [],zeros = [];
    if(!oDone){
        oFilter.forEach(o => {
            if(parseInt(lines[o][i]) == 1) ones.push(o);
            else zeros.push(o);
        });
        if(ones.length >= zeros.length) oFilter = [...ones];
        else oFilter = [...zeros];
        if(oFilter.length == 1) oDone = true;    
    }

    ones = [],zeros = [];
    if(!cDone){
        cFilter.forEach(c => {
            if(parseInt(lines[c][i]) == 1) ones.push(c);
            else zeros.push(c);
        });
        if(ones.length < zeros.length) cFilter = [...ones];
        else cFilter = [...zeros];
        if(cFilter.length == 1) cDone = true;
    }

}

let O = parseInt( lines[oFilter[0]],2 );
let C = parseInt( lines[cFilter[0]],2 );

console.log(C * O);