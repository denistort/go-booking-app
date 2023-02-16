((global) => {
    const validate = (formData) => {
        for (const formDatum of formData) {
            if (formDatum[1] === "") return false
        }
        return true
    }
    document.querySelector(".contact-form").addEventListener("submit", async (e) => {
        e.preventDefault()
        const formData = new FormData(document.querySelector(".contact-form"))
        if (!validate(formData)) {
            return
        }
        let data;
        try {
            const w = await fetch("/contact-us", {
                method: "post",
                body: formData
            });
            data = await w.json();
        } catch (err) {
            global.notie.alert({ type: 'error', text: 'Something went wrong', time: 2 })
        }
        if (data.ok === true) {
            global.notie.alert({ type: 'success', text: 'Your reservation has send', time: 2 })
            setTimeout(() => global.location.replace("/"), 2000)
        } else {
            global.notie.alert({ type: 'error', text: 'Something went wrong', time: 2 })
        }
        console.log(data)
    })
})(this)