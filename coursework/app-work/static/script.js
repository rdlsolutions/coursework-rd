document.addEventListener('DOMContentLoaded', function() {
    const calendarEl = document.getElementById('calendar');
    const calendar = new FullCalendar.Calendar(calendarEl, {
        initialView: 'dayGridMonth',
        locale: 'ua', // Встановлюємо українську локалізацію
        headerToolbar: {
            left: 'prev,next today',
            center: 'title',
            right: 'dayGridMonth,timeGridWeek,timeGridDay'
        },
        events: '/api/events', 
        
        // Обробник кліку по даті для заповнення форми
        dateClick: function(info) {
            // alert('Ви вибрали день: ' + info.dateStr);
            document.getElementById('event-start').value = info.dateStr + "T12:00";
            document.getElementById('event-end').value = info.dateStr + "T13:00";
        }
    });

    calendar.render();
    window.calendar = calendar;  // Зберігаємо екземпляр календаря у глобальну змінну для подальшого використання
});

// Функція для додавання нового події
function addNewEvent() {
    const title = document.getElementById('event-title').value;
    const start = document.getElementById('event-start').value;
    const end = document.getElementById('event-end').value;

    if (!title || !start || !end) {
        alert("Будь ласка, заповніть всі поля.");
        return;
    }

    const newEvent = {
        title: title,
        start_time: start + ":00Z", //  
        end_time: end + ":00Z"
    };

    fetch('/api/events', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(newEvent),
    })
    .then(response => response.json())
    .then(event => {
        console.log('Event added:', event);
        window.calendar.refetchEvents(); 
        document.getElementById('event-title').value = '';
    })
    .catch((error) => {
        console.error('Error adding event:', error);
        alert('Помилка при додаванні події.');
    });
}