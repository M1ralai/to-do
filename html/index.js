function get() {
  fetch("http://localhost:3000/index")
    .then((res) => res.json())
    .then((data) => {
      const container = document.getElementById("veriAlani");
      container.innerHTML = "";
        for (let i = 0; i < data.length; i++) {
          const div = document.createElement("div");
          div.innerHTML = `<h3>${data[i].id}: ${data[i].title}</h3>
                             <p>${data[i].description}</p>
                             <p>${data[i].creation_time}</p>
                             <p>${data[i].deadline}</p><hr>
                             <button onclick = "deleteAPI(${data[i].id})"> sil</button>`;

          container.appendChild(div);
        }
    })
    .catch((err) => console.error("Hata:", err));
}

function post() {
  fetch(
    "http://localhost:3000/index?title=Hello&desc=World&year=2025&month=12&day=31&hour=12&minute=30",
    {
      method: "POST",
      headers: {
        "Content-type": "application/json; charset=UTF-8",
      },
    }
  );
}

function deleteAPI(id) {
  fetch("http://localhost:3000/index?id=" + id, {
    method: "DELETE",
    headers: {
      "Content-type": "application/json; charset=UTF-8",
    },
  });
}
