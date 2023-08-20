const url = new URLSearchParams(window.location.search);

const _action = url.get("action");
const _status = url.get("status");

const SwalAlert = () => {
    Swal.fire({
        title: "Error!",
        text: "Do you want to continue",
        icon: "error",
        confirmButtonText: "Cool",
    });
};

(function() {
    if (_action == "insert") {
        SwalAlert();
    } else if (_action == "update") {
        SwalAlert();
    } else { // delete
        SwalAlert();
    }
})();