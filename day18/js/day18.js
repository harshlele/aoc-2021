const fs = require("fs");
const txt = fs.readFileSync("day18/input.txt");

const lines = txt.toString().split("\r\n");


const initNode = () => ({
  //for smaller multiples for eg 1000 or even 10000, 
  //it can produce 2 different answers in 2 different runs lmao
  //prob a better idea to use something like uuid i guess....
  id: parseInt(Math.random() * 10000000),       
  num: false,
  left: null,
  right: null,
  val: -1,
  parent: null,
});

//read line and convert to tree
function readInLine(line) {
  const root = initNode();
  let curr = root;
  let i = 1; 
  
  try{
    while (i < line.length - 1) {
      //for a starting bracket, start a new node and set it to current
      if(line[i] == "[") {
        if(curr.left == null) {
          curr.left = initNode();
          curr.left.parent = curr;
          curr = curr.left;
          i++;
          continue;
        }
        else if(curr.right == null) {
          curr.right = initNode();
          curr.right.parent = curr;
          curr = curr.right;
          i++;
          continue;
        }
        else {
          console.error("parsing error at i = ", i);
          break;
        }
      }
      if( !isNaN(parseInt(line[i])) ) {
        //read number until a , or a ] is encountered, this is a number node
        let no = parseInt(line[i]);
        i++;
        while(line[i] != "," && line[i] != "]") {
          no = no * 10 + parseInt(line[i])
          i++;
        }
        let node = initNode();
        node.num = true;
        node.val = no;
        node.parent = curr;
        if(curr.left == null) {
          curr.left = node;
        }
        else if(curr.right == null) {
          curr.right = node;
        }
        else{
          console.log("parsing error at i = ", i);
          break;
        }
        //console.log(`${no} at height ${getHeight(node)}`);
      }
  
      if(line[i] == "]") {
        curr = curr.parent;
      }
      
      if(line[i] == "]" || line[i] == ","){
        i++;
      }

    }  
  }
  catch(e) {
    console.log(e, "error at i = ", i)
  }
  return root;
}

//just print the tree for debugging
function printTree(node, list = []) {
  if(node.num) {
    return [`${node.val}`];
  }
  else {
    if(node.left) {
      return ["[",...list, ...printTree(node.left), ",", ...printTree(node.right), "]"];

    }
    else {
      console.log("owo")
    }
  }

}


//gets adjacent numbers of node using iterative dfs
function getAdjacentNums(root, nodeID) {
  let stack = [];
  let curr = root;

  let prevNum = null;
  let nodeFound = false;
  let nextNum = null;

  while(stack.length > 0 || curr != null) {
    if(curr != null) {
      stack.push(curr);
      curr = curr.left;
    }
    else {
      curr = stack.pop();
      
      if(curr.id == nodeID) {
        nodeFound = true;
      }

      if(curr.num == true && curr.parent.id != nodeID) {
        //store any number nodes
        if(!nodeFound) {
          prevNum = curr;
        }
        else {
          nextNum = curr;
          break;
        }
      
        
      }
      curr = curr.right;
    }
  }

  return [prevNum, nextNum];
}
//just gets the height
function getHeight(node) {
  let curr = node, height = 0;
  while(curr.parent != null) {
    curr = curr.parent;
    height++;
  }

  return height;
}

//the main thing...
//do iterative dfs from the root, if an explode/reduce action is to be done, 
//do that, then break out of the loop and check again from the beginning
//explodes always happen first, then splits, 
//and if a split produces an explode that explode happens before the remaining splits
function reduce(root){
  let stack = [];
  //check if explodes are to be checked for
  let explodesDone = false;
  while (true) {
    stack = [];
    let curr = root;  

    while (stack.length > 0 || curr != null) {
      if(curr != null) {
        stack.push(curr);
        curr = curr.left;
      }
      else {
        curr = stack.pop();
        if(!explodesDone) {
          let height = getHeight(curr);
          //do an explode
          if(height >= 5) {
            let adj = getAdjacentNums(root,curr.parent.id);
            explode(curr.parent, adj);
            //console.log(`after explode - ${printTree(root).join("")}`);
            break;
          }
        }
        else {
          //do a split
          if(curr.num && curr.val >= 10) {
            let parent = curr.parent;
            let newNode = split(curr.val);
            newNode.parent = parent;

            if(parent.left == curr) {
              parent.left = newNode;
            }
            else {
              parent.right = newNode;
            }
            
            //if a split increases height to more than 5, 
            //flag it so that node can be exploded in the next iteration
            if(getHeight(newNode.left) >= 5) {
              explodesDone = false;
            }
            break;
          }
        }
        
        curr = curr.right;
      }
    }
    
    if(stack.length == 0 && curr == null) {
      if(!explodesDone) {
        explodesDone = true;
        curr = root;
      }
      else break;
    }
  }
}

//split
//a number node split into inner node with num/2, num/2 as left/right children
function split(val) {
  let leftChild, rightChild;
  if(val % 2 == 0) {
    leftChild = rightChild = val/2;
  }
  else {
    leftChild = parseInt(val/2);
    rightChild = leftChild + 1;
  }

  let newNode = initNode();
  
  let newLeft = initNode();
  newLeft.num = true;
  newLeft.val = leftChild;
  newLeft.parent = newNode;

  let newRight = initNode();
  newRight.num = true;
  newRight.val = rightChild;
  newRight.parent = newNode;
  
  newNode.left = newLeft;
  newNode.right = newRight;

  return newNode;
}


//explode
//take an inner node, add left child to a predecessor number node(if exists),
//add right child to successor number node(if exists)
//replace the inner node with a 0
function explode(parNode, adj) {
  if(adj[0] && parNode.left) {
    adj[0].val += parNode.left.val;
  }
  if(adj[1] && parNode.right) {
    adj[1].val += parNode.right.val;
  }

  let newNode = initNode();
  newNode.val = 0;
  newNode.num = true;
  newNode.parent = parNode.parent;
  if(parNode.parent.left == parNode) {
    parNode.parent.left = newNode;
  }
  else {
    parNode.parent.right = newNode;
  }
  
}

//calculate magnitude using recursive dfs, mag(n) = 3 * mag(n.left) + 2 * mag(n.right)
function magnitude(root){
  
  if(root.num) {
    return root.val;
  }

  let totalSum = 0;

  if(root.left) {
    totalSum += 3 * magnitude(root.left);
  }

  if(root.right) {
    totalSum += 2 * magnitude(root.right);
  }
  
  return totalSum;
}


//read first input
let acc = readInLine(lines[0]);

for (let i = 1; i < lines.length; i++) {
  //read next line
  let curr = readInLine(lines[i]);

  //create new parent and append the accumulator and current tree to it
  let newRoot = initNode();
  newRoot.left = acc;
  newRoot.right = curr;
  acc.parent = newRoot;
  curr.parent = newRoot;
  //set new parent
  acc = newRoot;
  acc.parent = null;

  reduce(acc, acc);

}



console.log("-------------------------------------------------------");
console.log("final answer - ", printTree(acc).join(""));
console.log("magnitude - ", magnitude(acc));



