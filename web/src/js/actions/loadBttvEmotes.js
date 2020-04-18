import setBttvEmotes from "./setBttvEmotes";

export default function (channel) {
    return function (dispatch, getState) {
        return new Promise((resolve, reject) => {
            fetch("https://api.betterttv.net/2/emotes").then((response) => {
                if (response.status >= 200 && response.status < 300) {
                    return response
                } else {
                    var error = new Error(response.statusText)
                    error.response = response
                    throw error
                }
            }).then((response) => {
                return response.json();
            }).then((json) => {
                dispatch(setBttvEmotes(json))

                resolve();
            }).catch(err => {
                reject(err);
            });
        });
    };
}
