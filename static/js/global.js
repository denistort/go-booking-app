// Example starter JavaScript for disabling form submissions if there are invalid fields
(() => {
    'use strict'

    // Fetch all the forms we want to apply custom Bootstrap validation styles to
    const forms = document.querySelectorAll('.needs-validation')

    // Loop over them and prevent submission
    Array.from(forms).forEach(form => {
        form.addEventListener('submit', event => {
            if (!form.checkValidity()) {
                event.preventDefault()
                event.stopPropagation()
            }

            form.classList.add('was-validated')
        }, false)
    })
})()

const datesInputs = [...document.querySelectorAll('input[type="date"]')];

const startDayInput = document.querySelector("#date-from");
const endDayInput = document.querySelector("#date-to");

const DAYLIMITS = 5;

(() => {
    const options = { year: "numeric", month: "2-digit", day: "2-digit"};
    const day = 1000*60*60*24;
    const dateNow = Date.now() + day;
    const oneDayPlus = new Date(dateNow).toLocaleDateString("en", options);
    const splitted = oneDayPlus.split("/");
    const res = [splitted[2], splitted[0], splitted[1]].join("-")

    datesInputs.forEach(el => {
        el.min = res;
    })
    //
    endDayInput.disabled = true;

    startDayInput.addEventListener("change", (e) => {
        const valueOfChanged = e.target.value;
        if (valueOfChanged === "") {
            endDayInput.value = "";
            endDayInput.disabled = true;
            return
        }
        const valueNumeric = new Date(valueOfChanged).getTime();
        const res = valueNumeric + (day * DAYLIMITS);

        const plusLimitsInString = new Date(res).toLocaleDateString("en", options);
        const splitted = plusLimitsInString.split("/");
        endDayInput.disabled = false;
        endDayInput.min = [splitted[2], splitted[0], splitted[1]].join("-")
    })
})()

window.notie.alert({ type: 'success', text: 'Success!', time: 2 })

Swal.fire({
    title: 'Error!',
    text: 'Do you want to continue',
    icon: 'error',
    confirmButtonText: 'Cool'
})