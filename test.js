const TransicoesPermitidas = {
    StatusAguardando: [StatusCarregando,StatusCancelado], 
    StatusCarregando: [StatusFinalizado, StatusCancelado],
    StatusFinalizado: [], 
    StatusCancelado: [],
}

function PodeTrocarStatus(atual, novo) {
    const permitidos = TransicoesPermitidas[atual];
    return permitidos.includes(novo);
}