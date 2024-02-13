
document.addEventListener("DOMContentLoaded", function() {
    var currentPage = window.location.pathname;
    var links = document.querySelectorAll('a');

    links.forEach(function(link) {
        if (link.getAttribute('href') === currentPage) {
            link.style.fontWeight = 'bold';
        }
    });
});

function validateForm() {
    var checkInDate = document.getElementById('checkIn').value;
    var checkOutDate = document.getElementById('checkOut').value;

    if (checkOutDate <= checkInDate) {
        alert('Дата выезда должна быть после даты заезда');
        return false; 
    }

    return true; 
}
