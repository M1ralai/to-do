async function hello() {
    const veriAlani = document.getElementById('veriAlani'); // Hedef alanı unutma

    let id = 1;

    while (true) {
        try {
            const response = await fetch(`http://localhost:3000/index?id=${id}`);

            if (response.status === 404) {
                console.log(`ID ${id} bulunamadı. Döngü sonlandırılıyor.`);
                break; // 404 gelince dur
            }

            const data = await response.json();

            const div = document.createElement('div');
            div.innerHTML = `<h3>${data.ID}: ${data.title}</h3>
                             <p>${data.description}</p>
                             <p>${data.creation_time}</p>
                             <p>${data.deadline}</p><hr>`;
            veriAlani.appendChild(div);

            id++; // sıradaki ID'ye geç
        } catch (error) {
            console.error("Hata:", error);
            break; // Bağlantı hatası gibi durumlarda da döngüyü kes
        }
    }
}

