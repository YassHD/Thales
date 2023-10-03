function toggleMenu() {
  var navLinks = document.querySelector('.nav-links');
  var menuIcon = document.querySelector('.menu-icon');
  navLinks.classList.toggle('show');
  menuIcon.classList.toggle('active');
}
