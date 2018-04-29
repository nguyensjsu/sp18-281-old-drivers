var myIndex = 0;
displayImages();

/* The function displayImages() is used to create a slideshow
   of images in the recommendation section of the home page*/
function displayImages() {
    var i;
    var imgCount = document.getElementsByClassName("mySlides"); // extracts the number of images to be displayed
    for (i = 0; i < imgCount.length; i++) {
        imgCount[i].style.display = "none";
    }
    myIndex++;
    if (myIndex > imgCount.length) {myIndex = 1}
    imgCount[myIndex-1].style.display = "block";
    setTimeout(displayImages, 3000);// display each image after a 3 seconds gap
}

// When the user scrolls down 100px from the top of the document, show the button
window.onscroll = function() {
    scrollFunction()};

function scrollFunction() {
    if (document.body.scrollTop > 200 || document.documentElement.scrollTop > 200) {
        document.getElementById("toTopBtn").style.display = "block";
    } else {
        document.getElementById("toTopBtn").style.display = "none";
    }
}

// When the user clicks on the button, scroll to the top of the document
function topFunction() {
    document.body.scrollTop = 0; // For Chrome, Safari and Opera
    document.documentElement.scrollTop = 0; // For IE and Firefox
}
