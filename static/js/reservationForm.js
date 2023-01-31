(() => {
    const validate = (formData) => {
        for (const formDatum of formData) {
            if (formDatum[1] === "") return false
        }
        return true
    }
    document.querySelector("#reservation-form").addEventListener("submit", async (e) => {
        e.preventDefault()
        const formData = new FormData(document.querySelector("#reservation-form"))
        if (!validate(formData)) {
            return
        }
        const w = await fetch("/reservation", {
            method: "post",
            body: formData
        });
        const data = await w.json();
        console.log(data)
    })
})()