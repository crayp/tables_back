const rests = [{ id: 1, name: 'Zavodbar', stars: 5, description: '' },
               { id: 2, name: 'Траттория "Тепло"', stars: 4, description: ''},
               { id: 3, name: 'Наполи', stars: 5, description: ''},
               { id: 4, name: 'Ресторан "География"', stars: 5, description: '' },]
               
function initRests() {
  // if (!rests || rests.length === 0) {
  //   document.querySelector('#restr').innerHTML = '<p>Рестораны не найдены.</p>';
  //   return;
  // }

  console.log(rests)

  rests.forEach(rest => {
    const card = document.createElement('div');
    card.classList.add('row');

    card.innerHTML = `
        <div class="col-6 mainSectionItem">
          <div class="row h-75">
            <div class="col-12 d-flex justify-content-center align-items-center">
              <h3 class="text-center">${rest.name}</h3>
            </div>
          </div>
          <div class="row text-center h-25 d-flex justify-content-center align-items-center">
            <p>${rest.description || 'Описание не доступно.'}</p>
          </div>
        </div>
        <div class="col-6 mainSectionItem">
          <img src="/img/${rest.id}.webp" alt="${rest.name}" class="img-fluid" />
        </div>
    `;

    document.querySelector('#rests').appendChild(card);
  });
}

initRests()

// Получение всех пользователей
async function fetchUsers() {
    try {
      const response = await fetch("http://localhost:8080/api/users");
      if (!response.ok) {
        throw new Error("Ошибка при получении данных.");
      }
      const data = await response.json();
      const tableBody = document.querySelector("#users-table tbody");
      tableBody.innerHTML = "";

      data.forEach((user) => {
        const row = document.createElement("tr");
        row.innerHTML = `
                <td>${user.ID}</td>
                <td>${user.username || "—"}</td>
                <td>${user.email}</td>
                <td>${user.role || "—"}</td>
            `;
        tableBody.appendChild(row);
      });
    } catch (error) {
      console.error("Ошибка:", error);
      alert("Не удалось загрузить данные!");
    }
  }

  // Форма для добавления нового пользователя
  // document
  //   .querySelector("#add-user-form")
  //   .addEventListener("submit", async (e) => {
  //     e.preventDefault();

  //     const formData = new FormData(e.target);
  //     const data = {
  //       username: formData.get("username"),
  //       email: formData.get("email"),
  //       password: formData.get("password"),
  //       role: formData.get("role"),
  //     };

  //     try {
  //       const response = await fetch("http://localhost:8080/api/users", {
  //         method: "POST",
  //         headers: { "Content-Type": "application/json" },
  //         body: JSON.stringify(data),
  //       });

  //       if (response.ok) {
  //         alert("Пользователь успешно добавлен!");
  //         e.target.reset();
  //         //fetchUsers(); // Обновить список пользователей
  //       } else {
  //         throw new Error("Ошибка при добавлении пользователя.");
  //       }
  //     } catch (error) {
  //       console.error("Ошибка:", error);
  //       alert("Не удалось добавить пользователя.");
  //     }
  //   });

  // Загрузка данных при старте
  //document.addEventListener("DOMContentLoaded", fetchUsers);