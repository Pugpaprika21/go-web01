const urlParams = new URLSearchParams(window.location.search);

const _action = urlParams.get("action");
const _status = urlParams.get("status");

(function() {

    console.log([_action, _status]);

    //   Swal.fire({
    //     title: "Error!",
    //     text: "Do you want to continue",
    //     icon: "error",
    //     confirmButtonText: "Cool",
    //   });
})();

/* /?action=delete&status=Y */