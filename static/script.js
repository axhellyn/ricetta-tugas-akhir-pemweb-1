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
// searchBtn.addEventListener("click", (e) => {
//   e.preventDefault(); 
//   navButton.classList.toggle("active"); 
// });


// hamburger menu
const menu = document.querySelector ('.menu');
// ketika di klik
document.querySelector('#hamburger-menu').onclick = () => {
    menu.classList.toggle ('active');
};


// search function
document.addEventListener("DOMContentLoaded", () => {
    const searchInput = document.querySelector(".search-bar input");
    const recipeCards = document.querySelectorAll(".recipe-card");
    console.log(searchInput);

    searchInput.addEventListener("input", () => {
        const query = searchInput.value.toLowerCase();
        console.log(query);
        recipeCards.forEach(card => {
            const nama = card.getAttribute("data-name").toLowerCase();
            const details = card.getAttribute("data-details").toLowerCase();
            if (nama.includes(query) || details.includes(query)) {
                card.parentElement.style.display = ""; // Tampilkan
            } else {
                card.parentElement.style.display = "none";
            }
        });
    });
});