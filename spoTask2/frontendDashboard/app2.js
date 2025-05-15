let data

async function fetchSummary() {
            try {
                const res = await fetch("http://localhost:3000/summary", {
                    method: "GET",
                    headers: {
                        "Content-Type": "application/json"
                    }
                });

                if (!res.ok) throw new Error("Failed to fetch summary");

                 data = await res.json();

                console.log("Fetched summary data:", data);

                

                const arrStudentInfo = [...data];

            } catch (err) {
                console.error("Error fetching summary:", err);
            }
        }

        fetchSummary();


const div = document.querySelector("#addingStudendData")

data.forEach((student, index) => {
                    div.textContent = `Student #${index + 1} . "Name:", ${student.studentName}. "Roll No:", ${student.studentRollNo}. "Email:", ${student.studentEmail}`
                });


