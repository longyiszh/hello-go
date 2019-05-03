class Index {

  enablePetSelection = () => {
    let alinks = document.querySelectorAll("section.pet div.dContent a");
    let btn = document.querySelector("section.pet button.dropdown");
    for (let alink of alinks) {
      alink.addEventListener("click", () => {
        btn.innerHTML = alink.innerHTML;
      });
    }
  };

  enableQClassSelection = () => {
    let alinks = document.querySelectorAll("section.qclass div.dContent a");
    let btn = document.querySelector("section.qclass button.dropdown");
    for (let alink of alinks) {
      alink.addEventListener("click", () => {
        btn.innerHTML = alink.innerHTML;
      });
    }
  };

  main = () => {
    this.enablePetSelection();
    this.enableQClassSelection();
  };

  constructor() {
    this.main();
  }
}

const index = new Index();