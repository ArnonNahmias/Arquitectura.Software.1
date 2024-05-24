import React, { useState } from 'react';
import './Courses.scss';

const BoxCreator = () => {
  const [boxes, setBoxes] = useState([]);
  const [imageURL, setImageURL] = useState('');
  const [paragraph, setParagraph] = useState('');

  const addBox = () => {
    if (!imageURL ||!paragraph) return;
    const newBox = {
      id: Date.now(),
      imgSrc: imageURL,
      desc: paragraph,
    };
    setBoxes([...boxes, newBox]);
    setImageURL('');
    setParagraph('');
  };

  const removeBox = (id) => {
    setBoxes(boxes.filter((box) => box.id!== id));
  };

  const handleImageURLChange = (event) => {
    setImageURL(event.target.value);
  };

  const handleParagraphChange = (event) => {
    setParagraph(event.target.value);
  };

  return (
    <div className="container">
      <input
        type="text"
        value={imageURL}
        onChange={handleImageURLChange}
        placeholder="Enter image URL"
      />
      <textarea
        value={paragraph}
        onChange={handleParagraphChange}
        placeholder="Enter paragraph text"
      />
      <button onClick={addBox}>Add Box</button>
      <div className="boxes">
        {boxes.map((box) => (
          <div key={box.id} className="box">
            <img src={box.imgSrc} alt="box image" />
            <p>{box.desc}</p>
            <div
              className="remove"
              onClick={() => removeBox(box.id)}
            >
              X
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default BoxCreator;

