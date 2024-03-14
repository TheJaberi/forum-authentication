const join = document.querySelector(".join"),
  overlay = document.querySelector(".overlay"),
  closeBtn = document.querySelector(".overlay .close");

const create = document.querySelector(".createpost"),
  post = document.querySelector(".overlayposts"),
  closepost = document.querySelector(".overlayposts .close");

join.addEventListener("click", () => {
  overlay.classList.add("active");
});

closeBtn.addEventListener("click", () => { 
  overlay.classList.remove("active");
});

window.addEventListener("click", (e) => { 
  if (e.target === overlay){
    overlay.classList.remove("active");
  }
});

create.addEventListener("click", () => {
  post.classList.add("active");
});
closepost.addEventListener("click", () => {
  post.classList.remove("active");
});
window.addEventListener("click", (e) => { 
  if (e.target === post){
    post.classList.remove("active");
  }
});

const create2 = document.querySelector(".createcategory"),
catogary = document.querySelector(".overlaycatogaries"),
closecatogary = document.querySelector(".overlaycatogaries .close");

create2.addEventListener("click", () => {
  catogary.classList.add("active");
});
closecatogary.addEventListener("click", () => {
  catogary.classList.remove("active");
});
window.addEventListener("click", (e) => { 
  if (e.target === catogary){
    catogary.classList.remove("active");
  }
});

function togglePass() {
  var x = document.getElementById("passIn");
  var txt = document.getElementById("toggleTxt");
  if (x.type === "password") {
    x.type = "text";
    txt.textContent = "Hide";
  } else {
    x.type = "password";
    txt.textContent = "Show";
  }
}

function isEmail(email) {
  const emailRegex = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/;
  return emailRegex.test(email);
}

let loginemail = document.getElementById("loginemail");
if (isEmail(loginemail)) {
  alert("Login email incorrect")
}

let registeremail = document.getElementById("email");
if (isEmail(registeremail)) {
  alert("Register email incorrect")
}