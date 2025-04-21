document.getElementById('generate-btn').addEventListener('click', function () {
    const length = document.getElementById('length').value;

    fetch('/generate', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ length: parseInt(length) }),
    })
        .then(response => response.text())
        .then(result => {
            document.getElementById('result').textContent = result;
        })
        .catch(error => {
            console.error('Ошибка:', error);
            document.getElementById('result').textContent = 'Ошибка';
        });
});