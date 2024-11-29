// border bottom

window.addEventListener('scroll', () => {
    const header = document.querySelector("header")
    if (window.scrollY > 0) {
        header.classList.add("scrolled");
        console.log("discrol bg");
    } else {
        header.classList.remove("scrolled");
        console.log("engga");
    }
}); 



// search muncul
const searchBtn = document.getElementById("search");
const navButton = document.querySelector(".nav-button");

// Ketika tombol search diklik
searchBtn.addEventListener("click", (e) => {
  e.preventDefault(); 
  navButton.classList.toggle("active"); 
});


// hamburger menu
const menu = document.querySelector ('.menu');
// ketika di klik
document.querySelector('#hamburger-menu').onclick = () => {
    menu.classList.toggle ('active');
};
