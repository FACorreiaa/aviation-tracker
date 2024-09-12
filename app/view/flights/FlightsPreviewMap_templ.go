// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package flights

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import (
	"github.com/FACorreiaa/Aviation-tracker/app/models"
)

func previewMapContainer(data []models.LiveFlights) templ.ComponentScript {
	return templ.ComponentScript{
		Name: `__templ_previewMapContainer_5386`,
		Function: `function __templ_previewMapContainer_5386(data){const tileLayer = new ol.layer.Tile({
        source: new ol.source.OSM(),
    });


const map = new ol.Map({
    layers: [tileLayer],
    target: 'map',
    controls: [],
    view: new ol.View({
        center: [-11000000, 4600000],
        zoom: 1,
    }),
});

const style = new ol.style.Style({
    stroke: new ol.style.Stroke({
        color: '#EAE911',
        width: 3,
    }),
});

const markersSource = new ol.source.Vector();
const markersLayer = new ol.layer.Vector({
    source: markersSource,
});

map.addLayer(markersLayer);

for (let i = 0; i < data.length; i++) {
    const flight = data[i];

    const departureMarker = new ol.Feature({
        geometry: new ol.geom.Point(ol.proj.fromLonLat([
            parseFloat(flight.departure_longitude),
            parseFloat(flight.departure_latitude),
        ])),
        departure: flight.departure_airport,
        timezone: flight.departure.timezone
    });
    const departureMarkerStyle = new ol.style.Style({
        image: new ol.style.Icon({
            anchor: [0.5, 46],
            anchorXUnits: 'fraction',
            anchorYUnits: 'pixels',
            src: '../../../../static/icons/airplane-take-off.svg',
            scale: 0.5,
        }),
    });
    departureMarker.setStyle(departureMarkerStyle);
    markersSource.addFeature(departureMarker);

    const arrivalMarker = new ol.Feature({
        geometry: new ol.geom.Point(ol.proj.fromLonLat([
            parseFloat(flight.arrival_longitude),
            parseFloat(flight.arrival_latitude),
        ])),
        arrival: flight.arrival_airport,
        timezone: flight.arrival.timezone
    });
    const arrivalMarkerStyle = new ol.style.Style({
        image: new ol.style.Icon({
            anchor: [0.5, 46],
            anchorXUnits: 'fraction',
            anchorYUnits: 'pixels',
            src: '../../../static/icons/airplane-landing.svg',
            scale: 0.5,
        }),
    });
    arrivalMarker.setStyle(arrivalMarkerStyle);
    markersSource.addFeature(arrivalMarker);
}

const flightsSource = new ol.source.Vector({
    attributions: 'Flight data by ' + '<a href="https://openflights.org/data.html">OpenFlights</a>,',
    loader: function() {
        const features = [];

        for (let i = 0; i < data.length - 1; i++) {
            const from = data[i];
            const to = data[i + 1];


            const arcGenerator = new arc.GreatCircle({
                x: parseFloat(from.departure_longitude),
                y: parseFloat(from.departure_latitude),
            }, {
                x: parseFloat(to.departure_longitude),
                y: parseFloat(to.departure_latitude),
            });

            const arcLine = arcGenerator.Arc(100, {
                offset: 10
            });

            arcLine.geometries.forEach(function(geometry) {
                const line = new ol.geom.LineString(geometry.coords);
                line.transform('EPSG:4326', 'EPSG:3857');

                features.push(
                    new ol.Feature({
                        geometry: line,
                        departure: from.departure_airport,
                        arrival: to.arrival_airport,
                        finished: false,
                    })
                );
            });
        }

        addLater(features, 0);
        tileLayer.on('postrender', animateFlights);
    },
});


const flightsLayer = new ol.layer.Vector({
    source: flightsSource,
    style: function(feature) {
        if (feature.get('finished')) {
            return style;
        }
        return null;
    },
});

map.addLayer(flightsLayer);
const element = document.getElementById('popup');
const popup = new ol.Overlay({
    element: element,
    positioning: 'bottom-center',
    stopEvent: false
})

let popover;

function disposePopover() {
    if (popover) {
        popover.dispose();
        popover = undefined;
    } else {
        return
    }
}


const tippyButton = document.getElementById('popup');
tippy(tippyButton, {
    content: document.createElement('div'),
    interactive: true,
    trigger: 'click',
    placement: 'top',
    animation: 'scale',
    theme: 'translucent'
});

map.on('click', function(evt) {
    const feature = map.forEachFeatureAtPixel(evt.pixel, function(feature) {
        return feature;
    });
    disposePopover();
    if (!feature) {
        return;
    }
    popup.setPosition(evt.coordinate);

    // Show tooltip with departure or arrival information
    const contentDiv = document.createElement('div');
    if (feature.get('departure')) {
        contentDiv.innerHTML = ` + "`" + `<strong>Departure:</strong> ${feature.get('departure')}<br>
                                <p><strong>Timezone:</strong> ${feature.get('timezone')}<br></p>` + "`" + `;
    } else if (feature.get('arrival')) {
        contentDiv.innerHTML = ` + "`" + `<strong>Arrival:</strong> ${feature.get('arrival')}<br>
                                <p><strong>Timezone:</strong> ${feature.get('timezone')}<br></p>` + "`" + `;
    }
    tippyButton._tippy.setContent(contentDiv);
    tippyButton._tippy.show();
});

const pointsPerMs = 0.08;

function animateFlights(event) {
    const vectorContext = ol.render.getVectorContext(event);
    const frameState = event.frameState;
    vectorContext.setStyle(style);

    const features = flightsSource.getFeatures();
    for (let i = 0; i < features.length; i++) {
        const feature = features[i];
        if (!feature.get('finished')) {
            const coords = feature.getGeometry().getCoordinates();
            const elapsedTime = frameState.time - feature.get('start');
            if (elapsedTime >= 0) {
                const elapsedPoints = elapsedTime * pointsPerMs;

                if (elapsedPoints >= coords.length) {
                    feature.set('finished', true);
                }

                const maxIndex = Math.min(elapsedPoints, coords.length);
                const currentLine = new ol.geom.LineString(coords.slice(0, maxIndex));

                const worldWidth = ol.extent.getWidth(map.getView().getProjection().getExtent());
                const offset = Math.floor(map.getView().getCenter()[0] / worldWidth);

                currentLine.translate(offset * worldWidth, 0);
                vectorContext.drawGeometry(currentLine);
                currentLine.translate(worldWidth, 0);
                vectorContext.drawGeometry(currentLine);
            }
        }
    }
    map.render();
}

function addLater(features, timeout) {
    window.setTimeout(function() {
        let start = Date.now();
        features.forEach(function(feature) {
            feature.set('start', start);
            flightsSource.addFeature(feature);
            const duration = (feature.getGeometry().getCoordinates().length - 1) / pointsPerMs;
            start += duration;
        });
    }, timeout);
}

   document.getElementById('zoom-out').onclick = function () {
      const view = map.getView();
      const zoom = view.getZoom();
      view.setZoom(zoom - 1);
   };

   document.getElementById('zoom-in').onclick = function () {
      const view = map.getView();
      const zoom = view.getZoom();
      view.setZoom(zoom + 1);
   };

   map.on('dblclick', event => {
       // get the feature you clicked
       const feature = map.forEachFeatureAtPixel(event.pixel, (feature) => {
        return feature
       })
       if(feature instanceof ol.Feature){
         // Fit the feature geometry or extent based on the given map
         map.getView().fit(feature.getGeometry())
         // map.getView().fit(feature.getGeometry().getExtent())
       }
      })
}`,
		Call:       templ.SafeScript(`__templ_previewMapContainer_5386`, data),
		CallInline: templ.SafeScriptInline(`__templ_previewMapContainer_5386`, data),
	}
}

func FlightsLocationMap(data []models.LiveFlights) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<html><head><style scoped>\n                .map {\n                    width: 100%;\n                    height: 600px;\n                }\n                #map:focus {\n                    outline: #4A74A8 solid 0.15em;\n                }\n\t\t\t</style></head>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderScriptItems(ctx, templ_7745c5c3_Buffer, previewMapContainer(data))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<body onload=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 templ.ComponentScript = previewMapContainer(data)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ_7745c5c3_Var2.Call)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\"><div class=\"mb-5 text-left\"><button id=\"zoom-out\" class=\"btn btn-ghost mr-5\">Zoom out</button> <button id=\"zoom-in\" class=\"btn btn-ghost\">Zoom in</button></div><div id=\"map\" class=\"map\" tabindex=\"0\"><button aria-describedby=\"popup\" data-tippy-content=\"popup\" id=\"popup\"></button></div><script src=\"https://api.mapbox.com/mapbox.js/plugins/arc.js/v0.1.0/arc.js\"></script></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
