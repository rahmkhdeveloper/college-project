var ctx = document.getElementById("myChart").getContext("2d");
var newChart = new Chart(ctx, {
  type: "line",
  options: {
    elements: {
      point: {
        radius: 0,
      },
    },
  },
  data: {
    labels: [...Array(1000).fill(0)],
    datasets: [
      {
        data: [...Array(1000).fill(0)],
        backgroundColor: ["rgba(255, 99, 132, 0.2)"],
        borderColor: ["rgba(255, 99, 132, 1)"],
        borderWidth: 1,
      },
    ],
  },
});

var submit = document.getElementById("plot");
var cancel = document.getElementById("clear");
var initialRecords = true;
var count = 0;

submit.addEventListener("click", () => {
  axios({
    method: "post",
    url: "http://localhost:1323/",
    data: {
      n1: Number(document.getElementById("n1").value),
      n2: Number(document.getElementById("n2").value),
      d1: Number(document.getElementById("d1").value),
      d2: Number(document.getElementById("d2").value),
      th: Number(document.getElementById("th").value),
      n: Number(document.getElementById("n").value),
    },
  })
    .then((response) => {
      if (initialRecords) {
        newChart.data.datasets[0].data = response.data.TT11;
        newChart.data.datasets[0].label = "Plot " + ++count;
        updateConfigAsNewObject(newChart);
        initialRecords = false;
      } else {
        newChart.data.datasets.push({
          label: "Plot " + ++count,
          data: [...response.data.TT11],
          backgroundColor: [random_rgba()],
          borderColor: [random_rgba()],
          borderWidth: 1,
        });
        updateConfigAsNewObject(newChart);
      }
    })
    .catch((error) => console.log(error));
});

cancel.addEventListener("click", () => {
  newChart.data.datasets = [
    {
      data: [...Array(1000).fill(0)],
      backgroundColor: ["rgba(255, 99, 132, 0.2)"],
      borderColor: ["rgba(255, 99, 132, 1)"],
      borderWidth: 1,
    },
  ];
  newChart.data.labels = [...Array(1000).fill(0)];
  updateConfigAsNewObject(newChart);
  initialRecords = true;
  count = 0;
});

function updateConfigAsNewObject(chart) {
  chart.options = {
    responsive: true,
    elements: {
      point: {
        radius: 0,
      },
    },
    plugins: {
      legend: {
        display: true,
        position: "bottom",
        align: "center",
      },
      title: {
        display: true,
        text: "Graph Plot",
      },
    },
    scales: {
      x: {
        display: true,
        text: "lambda",
      },
      y: {
        display: true,
        text: "transmitivity",
      },
    },
  };
  chart.update();
}

function random_rgba() {
  var o = Math.round,
    r = Math.random,
    s = 255;
  return (
    "rgba(" +
    o(r() * s) +
    "," +
    o(r() * s) +
    "," +
    o(r() * s) +
    "," +
    r().toFixed(1) +
    ")"
  );
}
