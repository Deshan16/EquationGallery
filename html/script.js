const navContainer = document.createElement("div");
navContainer.classList.add("nav-items");

const links = [
    ["යාන්ත්‍ර විද්‍යාව", "file1.html"],
    ["දෝලන හා තරංග", "file2.html"],
    ["තාපය", "file3.html"],
    ["පදාර්ථයේ යාන්ත්‍රික ගුණ", "file4.html"],
    ["ධාරා විද්‍යුතය", "file5.html"]
];

let currentPage = window.location.pathname.split('/').pop();

if (currentPage === '') currentPage = "file1.html";

links.forEach(([name, fileName]) => {
    const a = document.createElement("a");
    a.textContent = name;
    a.href = `/page/${fileName}`;
    
    if (currentPage === fileName) a.classList.add("selected");
    
    navContainer.appendChild(a);
});

document.querySelector(".nav-bar").appendChild(navContainer);